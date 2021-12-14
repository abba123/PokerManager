package kafka

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func write() {
	kafkaURL := "localhost:9092"
	topic := "quickstart-events"

	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Value: []byte(fmt.Sprintf("name-%d", i)),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}
		time.Sleep(1 * time.Second)
	}
}

func read() {
	kafkaURL := "localhost:9092"
	topic := "quickstart-events"

	reader := getKafkaReader(kafkaURL, topic)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic: %v partition: %v offset: %v value: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
	}
}

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 0,    // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
