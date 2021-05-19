package handler

import (
	"context"
	"github.com/asim/go-micro/v3/broker"
	"log"
)

type DemoServiceHandler struct {
	PubSub broker.Broker
}

func ProcessEvent(ctx context.Context, event *broker.Message) error {
	log.Println("Got alert:", event)
	return nil
}

