package main

import (
	"context"
	kafka_handler "project/kafka-handler"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go kafka_handler.Produce(ctx)
	kafka_handler.Consumer(ctx)
}
