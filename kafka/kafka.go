package kafka

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"net"
	"strconv"
)

type KafkaHolder struct {
	writer      *kafka.Writer
	initialized bool
}

func (k *KafkaHolder) InitWriter(endpoint string, topic string) error {
	k.writer = &kafka.Writer{
		Addr:     kafka.TCP(endpoint),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	k.initialized = true
	return nil
}

func (k *KafkaHolder) Write(messagesBatch []string) error {
	var messages []kafka.Message

	for i := 0; i < len(messagesBatch); i++ {
		bytesPayload := []byte(messagesBatch[i])
		kafkaPayload := kafka.Message{Value: bytesPayload}
		messages = append(messages, kafkaPayload)
	}

	if k.writer != nil && k.initialized == true {
		err := k.writer.WriteMessages(context.Background(), messages...)
		if err != nil {
			k.initialized = false
			return err
		}
	} else {
		log.Println("kafka writer not initialized")
		return errors.New("kafka writer not initialized")
	}
	return nil
}

func (k *KafkaHolder) Close() error {
	k.initialized = false
	k.writer = nil
	err := k.writer.Close()
	return err
}

func CreateTopic(endpoint string, topic string, ) error {

	conn, err := kafka.Dial("tcp", endpoint)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func DeleteTopic(endpoint string, topic string) error {
	conn, err := kafka.Dial("tcp", endpoint)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	err = controllerConn.DeleteTopics(topic)
	return err
}
