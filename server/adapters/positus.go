package adapters

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Positus struct {
	Url   string
	Token string
}

func (p Positus) GetUrl() string {
	return os.Getenv("POSITUS_URL")
}

func (p Positus) GetToken() string {
	return os.Getenv("POSITUS_TOKEN")
}

func CreatePositusObject(url string, token string) Positus {
	var positusObject Positus
	positusObject.Url = url
	positusObject.Token = token
	return positusObject
}

func (p Positus) SendMessage(widReceiver string, message string) error {
	payload := strings.NewReader("{\n" +
		"	\"to\": \"" + widReceiver + "\",\n" +
		"	\"type\": \"text\",\n" +
		"	\"text\": {\n" +
		"		\"body\": \"" + string([]rune(message)) + "\"\n" +
		"	}" +
		"}")
	fmt.Println(payload)

	url := p.GetUrl()

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

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
