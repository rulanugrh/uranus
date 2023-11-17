package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Payload struct {
	Course   string
	Category string
	User     string
	Payment  string
}

func SendWebhook(data interface{}, url string, webhookID string) error {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		log.Println("Can't marshalling data in json")
	}

	req, errRequest := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	if errRequest != nil {
		log.Println("Cant request to this url")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		log.Println("Cant request client")
	}

	defer func(Body io.ReadCloser) {
		if errBody := Body.Close(); errBody != nil {
			log.Println("cant reader by body request")
		}
	}(resp.Body)

	status := "failed"
	if resp.StatusCode == http.StatusOK {
		status = "delivered"
	}

	log.Println(status)

	if status == "failed" {
		return errors.New(status)
	}

	return nil
}
