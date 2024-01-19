package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Create a new Kafka producer
	//producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer producer.Close()
	//
	//// Create a new Kafka topic
	//topic := "kafka.learning.orders"
	//fmt.Println("Topic Created ", topic)
	letterNotes := []string{"A", "B", "C", "D", "E", "F", "G"}
	for i := 0; i < 100; i++ {
		//message := &sarama.ProducerMessage{
		//	Topic: topic,
		//	Value: sarama.StringEncoder(fmt.Sprintf("Student Number : ! %d Letter Score :  %s", rand.Int31n(1000), letterNotes[rand.Int31n(int32(len(letterNotes)))])),
		//}
		fmt.Println("Student Number : ", rand.Int31n(1000), "Letter Score : ", letterNotes[rand.Int31n(int32(len(letterNotes)))])
		//_, _, err = producer.SendMessage(message)
		//fmt.Println("Message Generated ", message.Value)
		//time.Sleep(1 * time.Second)
		//if err != nil {
		//	log.Fatal(err)
		//	//continue
		//}

	}

	//}()

}
