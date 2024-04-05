package main

import (
	"fmt"
	"testing"
	"time"

	kafkaConfig "go-kafka-mqtt/kafka/config"
	mqttConfig "go-kafka-mqtt/mqtt/config"
	mqttPublisher "go-kafka-mqtt/mqtt/publisher"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"
)

func TestMQTTConnection(t *testing.T) {
	wg := sync.WaitGroup{}
	mqttClient := mqttConfig.MQTTClient{Wg: &wg}
	client := mqttClient.Config()
	// O cliente deve estar conectado, podemos verificar isso usando IsConnected
	if !client.IsConnected() {
		t.Fatal("Failed to connect to MQTT broker")
	}
	client.Disconnect(0)
	fmt.Println("Successfully connected to MQTT broker")
}

func TestKafkaConnection(t *testing.T) {
	conf := kafkaConfig.ReadConfig("../client.properties")
	conf["group.id"] = "test-group"
	conf["auto.offset.reset"] = "earliest"

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		t.Fatal("Failed to create Kafka consumer:", err)
	}

	err = consumer.Close()
	if err != nil {
		t.Fatal("Failed to close Kafka consumer:", err)
	}

	fmt.Println("Successfully connected to Kafka broker")
}

func TestMQTTToKafkaIntegration(t *testing.T) {
	// Configurar e conectar o cliente MQTT
	wg := sync.WaitGroup{}
	mqttClient := mqttConfig.MQTTClient{Wg: &wg}
	client := mqttClient.Config()
	defer client.Disconnect(250)

	// Publicar 10 mensagens via MQTT
	go func() {
		mqttPublisher.Publish(client)
	}()

	// Configurar o consumidor Kafka
	conf := kafkaConfig.ReadConfig("../client.properties")
	conf["group.id"] = "test-group"
	conf["auto.offset.reset"] = "earliest"
	topic := "movies"

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		t.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		t.Fatalf("Failed to subscribe to Kafka topic: %v", err)
	}

	// Consumir mensagens do Kafka
	receivedMessages := make(map[string]bool)
	timeout := time.After(30 * time.Second)
	for len(receivedMessages) < 10 {
		select {
		case <-timeout:
			t.Fatal("Timeout reached before all messages were consumed")
		default:
			ev := consumer.Poll(10)
			switch e := ev.(type) {
			case *kafka.Message:
				receivedMessages[string(e.Value)] = true
			case kafka.Error:
				t.Fatalf("Kafka error: %v", e)
			}
		}
	}

	// Verificar se todas as mensagens foram recebidas
	assert.Equal(t, 10, len(receivedMessages), "Not all messages were received")

	fmt.Println("All messages were successfully published to MQTT and consumed from Kafka")
}
