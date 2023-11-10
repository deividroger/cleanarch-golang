package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/deividroger/cleanarch-golang/pkg/events"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		viper.GetString("RABBITMQ_ORDER_CREATED_EXCHANGE"), // exchange
		viper.GetString("RABBITMQ_KEY_ORDER_CREATED"),      // key name
		false,       // mandatory
		false,       // immediate
		msgRabbitmq, // message to publish
	)
}
