package main

import (
	"chatgpt-article/internal/handlers"
	"chatgpt-article/internal/kafka"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

const (
	topic = "ai-article-topic"
	consumerGroup  = "ai-article-group"
)

var address = []string{"localhost:9091"}

func main() {
	h := handlers.NewHandler()

	c, err := kafka.NewConsumer(h, address, topic, consumerGroup)
	if err != nil {
		logrus.Fatal(err)
	}

	go c.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	logrus.Fatal(c.Stop())
}