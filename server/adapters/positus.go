package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Text struct {
	Body string `json:"body"`
}
type Positus struct {
	To   string `json:"to"`
	Type string `json:"type"`
	Text Text   `json:"text"`
}

func (p Positus) GetUrl() string {
	return os.Getenv("POSITUS_URL")
}

func (p Positus) GetToken() string {
	return os.Getenv("POSITUS_TOKEN")
}

func (p Positus) SendMessage(widReceiver string, message string) error {
	Message := &Positus{To: widReceiver, Type: "text", Text: Text{Body: message}}
	b, _ := json.Marshal(Message)

	fmt.Println(b)

	url := p.GetUrl()

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+p.GetToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	return err
}

// {
//   "contacts": [ {
//     "profile": {
//         "name": "Kerry Fisher"
//     },
//     "wa_id": "16315551234"
//   } ],
//   "messages":[{
//     "from": "16315551234",
//     "id": "ABGGFlA5FpafAgo6tHcNmNjXmuSf",
//     "timestamp": "1518694235",
//     "text": {
//       "body": "Hello this is an answer"
//     },
//     "type": "text"
//   }]

type ResponseText struct {
	Body string `json:"body"`
}
type ResponseMessage struct {
	From      string       `json:"from"`
	ID        string       `json:"id"`
	Timestamp string       `json:"timestamp"`
	Text      ResponseText `json:"text"`
	Type      string       `json:"type"`
}
type ResponseContact struct {
	WidSender string `json:"wa_id"`
}
type ResponseType struct {
	Contacts []ResponseContact `json:"contacts"`
	Messages []ResponseMessage `json:"messages"`
}
