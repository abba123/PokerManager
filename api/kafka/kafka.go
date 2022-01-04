package kafka

import (
	"context"
	"fmt"
	"log"
	"poker/api/model"
	"poker/poker"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

//const kafkaURL string = "ec2-3-131-38-31.us-east-2.compute.amazonaws.com:9092"

const topic string = "pokerHand"

func NewKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func NewKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := []string{kafkaURL}
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		GroupID:     "consumer",
		Topic:       topic,
		MinBytes:    0,    // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: -1,
	})
}

func KafkaWrite(data []byte, username []byte) {
	viper.AutomaticEnv()
	kafkaURL := viper.GetString("DATABASE") + ":9092"
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

	viper.AutomaticEnv()
	kafkaURL := viper.GetString("DATABASE") + ":9092"
	reader := NewKafkaReader(kafkaURL, topic)

	defer reader.Close()

	fmt.Println("start consuming ... !!")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		tables := poker.Parsefile(string(m.Value))
		model.InsertHandDB("pokerdb",string(m.Key), tables)
		model.RemoveKeyRedis(string(m.Key))

	}
}
