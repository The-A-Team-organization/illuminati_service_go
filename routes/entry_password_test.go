package routes

import (
	"encoding/json"
	"fmt"
	mock_utils "illuminati/go/microservice/utils/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)


func Test_GetRandomWord(t *testing.T) {
    word := getRandomWord()
    if strings.TrimSpace(word) == "" {
        t.Error("Expected non-empty word")
    }
}

func Test_GetAppParticipants(t *testing.T) {
    mockParticipantsAnswer :=  func(w http.ResponseWriter, r *http.Request) {
        type data struct {
	 	Participants []string `json:"participants"`}
        var example data 
        example.Participants = []string{"test1@gmail.com", "test123@gmail.com"}
        dataSend, _  := json.Marshal(example)
        fmt.Println(dataSend)
        w.Header().Set("Content-Type", "application/json")
        w.Write(dataSend)
    }
    
    server := httptest.NewServer(http.HandlerFunc(mockParticipantsAnswer))
    defer server.Close()

    var (
		ctrl            = gomock.NewController(t)
		mockEmailSender = mock_utils.NewMockEmailSender(gomock.NewController(t))
		service         = NewEntryPasswordService(mockEmailSender, server.URL)
	)
    defer ctrl.Finish()

    resp, _  := service.getAppParticipants()
    if resp == nil {
        t.Errorf("Expected post@pon.com, post1@pon.com, post2@pon.com, got %s", resp)
    }
}



func Test_sendEntryPasswordEmail(t *testing.T) {

    word := getRandomWord()

    text := fmt.Sprintf(`Hello,

		You're subscribed to our Latin vocabulary service, and today’s Word of the Day is %s.

		Learn its meaning, usage, and examples to expand your vocabulary!

		Happy learning
		— The Latin Words Team
		`,
	word)


    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
    defer server.Close()

    var (
		ctrl            = gomock.NewController(t)
		mockEmailSender = mock_utils.NewMockEmailSender(gomock.NewController(t))
		service         = NewEntryPasswordService(mockEmailSender, server.URL)
        participants = []string{"test1@gmail.com", "test123@gmail.com"}
	)
	defer ctrl.Finish()



	mockEmailSender.EXPECT().SendEmail("Word of the Day", text, participants).Return(nil)

	err := service.sendEntryPasswordEmail(word,participants)
	assert.NoError(t, err)
    
}


func Test_getNewEntryPassword(t *testing.T) {

    mockParticipantsAnswer :=  func(w http.ResponseWriter, r *http.Request) {
        type data struct {
	 	Participants []string `json:"participants"`}
        var example data 
        example.Participants = []string{"test1@gmail.com", "test123@gmail.com"}
        dataSend, _  := json.Marshal(example)
        fmt.Println(dataSend)
        w.Header().Set("Content-Type", "application/json")
        w.Write(dataSend)
    }
    
    back_server := httptest.NewServer(http.HandlerFunc(mockParticipantsAnswer))

    word := getRandomWord()

    text := fmt.Sprintf(`Hello,

		You're subscribed to our Latin vocabulary service, and today’s Word of the Day is %s.

		Learn its meaning, usage, and examples to expand your vocabulary!

		Happy learning
		— The Latin Words Team
		`,
	word)

  

    var (
		ctrl            = gomock.NewController(t)
		mockEmailSender = mock_utils.NewMockEmailSender(gomock.NewController(t))
		service         = NewEntryPasswordService(mockEmailSender, back_server.URL)
        participants = []string{"test1@gmail.com", "test123@gmail.com"}
	)
	defer ctrl.Finish()



	mockEmailSender.EXPECT().SendEmail("Word of the Day", text, participants).Return(nil)
    

    server := httptest.NewServer(http.HandlerFunc(service.getNewEntryPassword))
    defer server.Close()

    response, _ := http.Get(server.URL)

	assert.Equal(t, response.Status, "200 OK")
    
}

