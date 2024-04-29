package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	servers := []string{"localhost:9092"}
	consumer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	partition, err := consumer.ConsumePartition("bondhi", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partition.Close()

	fmt.Println("Consumer start.")
	for {
		select {
		case err := <-partition.Errors():
			fmt.Println(err)
		case msg := <-partition.Messages():
			println(string(msg.Value))
		}
	}
}
