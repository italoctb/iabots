package api

import (
	"app/server/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Adapter models.Positus
}

func (api Api) SendTextMessage(widReceiver string, message string, method string) error {
	url := api.Adapter.Url

	payload := api.Adapter.TextMessage(widReceiver, message, method)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+api.Adapter.Token)

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
