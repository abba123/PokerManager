package kafka

import (
	"context"
	"fmt"
	"log"
	"poker/api/model"
	"poker/poker"
	"strings"

	"github.com/segmentio/kafka-go"
)

//const kafkaURL string = "ec2-3-131-38-31.us-east-2.compute.amazonaws.com:9092"
const kafkaURL string = "localhost:9092"
const topic string = "pokerHand" 

func NewKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func NewKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		MinBytes:    0,    // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: -1,
	})
}

func KafkaWrite(data []byte, username []byte) {
	writer := NewKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	msg := kafka.Message{
		Value: data,
		Key:   username,
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	}

}

func KafkaRead() {
	topic := "pokerHand"

	reader := NewKafkaReader(kafkaURL, topic)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	reader.SetOffset(-1)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		table := poker.Parsefile(string(m.Key), string(m.Value))
		model.InsertHandDB(table)
		for _,p := range table[0].Player{
			model.RemoveKeyRedis(p.Name)
		}
	}
}
