package kafka

import (
	"fmt"

	kafka "github.com/IBM/sarama"
)

func InitProducer() (kafka.SyncProducer, error) {
	brokersUrl := []string{"localhost:9092"}
	config := kafka.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = kafka.WaitForAll
	config.Producer.Retry.Max = 5
	conn, err := kafka.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}


func Produce(topic string, message []byte) error {
	producer, err := InitProducer()
	if err != nil {
			return err
	}
	defer producer.Close()
	msg := &kafka.ProducerMessage{
			Topic: topic,
			Value: kafka.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
			return err
	}
	fmt.Println( topic, partition, offset)
	return nil
}