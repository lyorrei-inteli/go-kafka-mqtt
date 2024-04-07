package main

import (
	"fmt"
	kafkaConsumer "go-kafka-mqtt/kafka/consumer"
	"sync"

	// kafkaPublisher "go-kafka-mqtt/kafka/publisher"
	mqttConfig "go-kafka-mqtt/mqtt/config"
	mqttPublisher "go-kafka-mqtt/mqtt/publisher"
)

func main() {
	wg := sync.WaitGroup{}
	// kafkaPublisher.Publish()
	mqttClient := mqttConfig.MQTTClient{Wg: &wg}
	client := mqttClient.Config()

	wg.Wait()

	fmt.Println("Starting Kafka-MQTT Client")
	go mqttPublisher.Publish(client)
	go kafkaConsumer.Consume()
	select {}
}
