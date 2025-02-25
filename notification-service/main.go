package main

import (
	"github.com/telman03/go-microservices/notification-service/config"
	"github.com/telman03/go-microservices/notification-service/kafka"
)

func main() {
	config.LoadEnv()
	kafka.ConsumeMessages()
}