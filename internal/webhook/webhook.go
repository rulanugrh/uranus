package webhook

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type WebhookPayload struct {
	Url       string `json:"url"`
	WebhookId string `json:"webhookId"`
	Data      struct {
		Id       string `json:"id"`
		User     string `json:"user"`
		Course   string `json:"course"`
		Category string `json:"category"`
		Price    int    `json:"price"`
		Payment  string `json:"payment"`
	} `json:"data"`
}

func Subscribe(ctx context.Context, client *redis.Client, webhookQueue chan WebhookPayload) error {
	pubSub := client.Subscribe(ctx, "payments")
	defer func(pubSub *redis.PubSub) {
		if err := pubSub.Close(); err != nil {
			log.Println("Error closing PubSub:", err)
		}
	}(pubSub)
	var payload WebhookPayload

	for {
		msg, err := pubSub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("Error recive message")
			return err
		}

		err = json.Unmarshal([]byte(msg.Payload), &payload)
		if err != nil {
			log.Println("Error unmarshalling payload:", err)
			continue
		}

		webhookQueue <- payload
	}
}
