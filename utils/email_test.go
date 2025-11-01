package utils

import (
	mock_utils "illuminati/go/microservice/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)


func Test_SendEmail(t *testing.T) {

	var (
		ctrl            = gomock.NewController(t)
		mockEmailSender = mock_utils.NewMockEmailSender(gomock.NewController(t))
		topic  			= "New Letter"
		text 			= "Here is new letter"
		targetEmails 	= []string{"test1@gmail.com", "test123@gmail.com"}
		
	)
	defer ctrl.Finish()

	mockEmailSender.EXPECT().SendEmail(topic, text, targetEmails).Return(nil)

	err := mockEmailSender.SendEmail(topic, text, targetEmails)
	assert.NoError(t, err)
}


