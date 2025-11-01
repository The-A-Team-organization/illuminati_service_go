package routes

import (
	"bytes"
	"encoding/json"
	mock_utils "illuminati/go/microservice/utils/mocks"
	"net/http"
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
