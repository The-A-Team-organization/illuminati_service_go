package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	mock_utils "illuminati/go/microservice/utils/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_SendLetterEmail(t *testing.T) {
	var (
		ctrl            = gomock.NewController(t)
		mockEmailSender = mock_utils.NewMockEmailSender(gomock.NewController(t))
		service         = NewLetterService(mockEmailSender)
		letter          = Letter{
			Topic:        "New Letter",
			Text:         "Here is new letter",
			TargetEmails: []string{"test1@gmail.com", "test123@gmail.com"},
		}
	)
	defer ctrl.Finish()

	jsonLetter, _ := json.Marshal(letter)

	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonLetter))

	mockEmailSender.EXPECT().SendEmail(letter.Topic, letter.Text, letter.TargetEmails).Return(nil)

	err := service.SendLetterEmail(request)
	assert.NoError(t, err)
}


func Test_SendLetterEmailWhileRequestBodyIsNil(t *testing.T) {
	var (
		ctrl            = gomock.NewController(t)
		mockEmailSender = mock_utils.NewMockEmailSender(gomock.NewController(t))
		service         = NewLetterService(mockEmailSender)
	)
	defer ctrl.Finish()


	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(nil))


	err := service.SendLetterEmail(request)
	assert.Error(t, err)
}

func Test_PostLetter(t *testing.T) {

	
	var (
		ctrl            = gomock.NewController(t)
		mockEmailSender = mock_utils.NewMockEmailSender(gomock.NewController(t))
		service         = NewLetterService(mockEmailSender)
		letter          = Letter{
			Topic:        "New Letter",
			Text:         "Here is new letter",
			TargetEmails: []string{"test1@gmail.com", "test123@gmail.com"},
		}
	)
	defer ctrl.Finish()

	jsonLetter, _ := json.Marshal(letter)

	mockEmailSender.EXPECT().SendEmail(letter.Topic, letter.Text, letter.TargetEmails).Return(nil)

	server := httptest.NewServer(http.HandlerFunc(service.PostLetter))
	defer server.Close()

	response, _ := http.Post(server.URL, "application/json" ,bytes.NewBuffer(jsonLetter))



	assert.Equal(t, response.Status, "202 Accepted")

}

func Test_PostLetterWhileErrIsNotNil(t *testing.T) {

	
	var (
		ctrl            = gomock.NewController(t)
		mockEmailSender = mock_utils.NewMockEmailSender(gomock.NewController(t))
		service         = NewLetterService(mockEmailSender)
		letter          = Letter{
			Topic:        "New Letter",
			Text:         "Here is new letter",
			TargetEmails: []string{},
		}
	)
	defer ctrl.Finish()

	jsonLetter, _ := json.Marshal(letter)

	mockEmailSender.EXPECT().SendEmail(letter.Topic, letter.Text, letter.TargetEmails).Return(errors.New("receivers cant`t be null"))

	server := httptest.NewServer(http.HandlerFunc(service.PostLetter))
	defer server.Close()

	response, _ := http.Post(server.URL, "application/json" ,bytes.NewBuffer(jsonLetter))



	assert.NotEqual(t, response.Status, "202 Accepted")

}
