package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Consume messages from the topic
	// Create a new Kafka consumer
	topic := "kafka.learning.orders"
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil { // hata varsa
		log.Fatal(err)
	}
	// defer nedir ?
	// defer, bir fonksiyonun sonunda çalıştırılması gereken bir işlevi belirtmek için kullanılır.

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}
	defer func(partitionConsumer sarama.PartitionConsumer) {
		err := partitionConsumer.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(partitionConsumer)

	// Create a new signal receiver
	// Signal receiver is used to stop the consumer loop
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Start consuming messages in a separate goroutine
	// The loop will stop when the signal receiver gets an interrupt signal
	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("Received message: %s\n", string(msg.Value))
			case <-signals:
				return
			}
		}
	}()

	// Wait for an interrupt signal to stop the consumer
	<-signals
	fmt.Println("Consumer stopped")
}
