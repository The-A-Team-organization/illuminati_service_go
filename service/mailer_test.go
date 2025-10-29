package service

import (
	"fmt"
	"strings"
	"testing"

	"github.com/h2non/gock"
)

func Test_GetRandomWord(t *testing.T) {
    word := GetRandomWord()
    if strings.TrimSpace(word) == "" {
        t.Error("Expected non-empty word")
    }
}

// func Test_SendEmail(t *testing.T) {

//     emails := []string{"post@pon.com"}
//     sendTests := []struct {
//         word string
//         emails []string
//     }{
//         {"lorem", nil},
//         {"", emails},
//     }
    
// }

func Test_GetAppParticipants(t *testing.T) {

    //jsontest := `"participants": ["test1@gmail.com", "test123@gmail.com"]`

    // httpmock.Activate()
    // defer httpmock.DeactivateAndReset()
    
    // httpmock.RegisterResponder(
    //     http.MethodGet, 
    //     "https://example.com",
    //     func(req *http.Request) (*http.Response, error) {
    //        resp, err := httpmock.NewJsonResponse(200, httpmock.File("test_body.json"))
    //        fmt.Print(resp)
    //        return resp, err
    //     },
    // )

    // 

	// json.NewDecoder(value.Body).Decode(&data);

    // fmt.Println()
    // fmt.Println()
    // fmt.Println(value.Body)
    // fmt.Println(data.participants)
    // fmt.Println()
    // fmt.Println()

    type data struct {
	 	Participants []string `json:"participants"`
	}

    expected := &data{
        Participants:[]string{"test1@gmail.com", "test123@gmail.com"} ,
    }
    
   // var resp data

    // {
    //     Participants : []s{"test1@gmail.com", "test123@gmail.com"},
    // }

    gock.New("https://example.com").
       Get("/").
       Reply(200).
       JSON(expected)

    value, _ := GetAppParticipants("https://example.com/")


    // fmt.Printf("%s", value.Body)

    // err = json.NewDecoder(value.Body).Decode(&resp)
    // if err != nil {
    //         panic(err)
    // }
    // fmt.Println()
    fmt.Println(value)
    // if data.Participants == nil {
    //     t.Errorf("Expected post@pon.com, post1@pon.com, post2@pon.com, got %s, error: %s", data.Participants , err)
    // }
 
}
