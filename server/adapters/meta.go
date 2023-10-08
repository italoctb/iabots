package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TextMeta struct {
	Body string `json:"body"`
}
type Meta struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             Text   `json:"text"`
}

func (m Meta) GetUrl() string {
	return os.Getenv("META_URL")
}

func (m Meta) GetToken() string {
	return os.Getenv("META_TOKEN")
}

func (m Meta) SendMessage(widReceiver string, message string) error {
	Message := &Meta{To: "5585997112838", Type: "text", Text: Text{Body: message}, MessagingProduct: "whatsapp"}
	b, _ := json.Marshal(Message)

	//fmt.Println(b)

	url := m.GetUrl()

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+m.GetToken())

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

type MetaResponseMessage struct {
	From      string       `json:"from"`
	ID        string       `json:"id"`
	Timestamp string       `json:"timestamp"`
	Text      ResponseText `json:"text"`
	Type      string       `json:"type"`
}
type MetaResponseContact struct {
	Profile     MetaProfileType `json:"profile"`
	WidReceiver string          `json:"wa_id"`
}
type MetaResponseType struct {
	MessagingProduct string                `json:"messaging_product"`
	Metadata         MetaDataType          `json:"metadata"`
	Contacts         []MetaResponseContact `json:"contacts"`
	Messages         []MetaResponseMessage `json:"messages"`
}

type MetaDataType struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberId      string `json:"phone_number_id"`
}

type MetaProfileType struct {
	Name string `json:"name"`
}

type MetaResponseObject struct {
	Object string            `json:"object"`
	Entry  []MetaObjectEntry `json:"entry"`
}

type MetaObjectEntry struct {
	Id      string            `json:"id"`
	Changes []MetaEntryChange `json:"changes"`
}

type MetaEntryChange struct {
	Field string           `json:"field"`
	Value MetaResponseType `json:"value"`
}
