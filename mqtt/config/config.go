package mqtt

import (
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	Wg *sync.WaitGroup
}

// upon connection to the client, this is called
func (s *MQTTClient) connectHandler(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker")
	defer s.Wg.Done()
}

// this is called when the connection to the client is lost, it prints "Connection lost" and the corresponding error
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost to MQTT broker: %v", err)
}

func (s *MQTTClient) Config() mqtt.Client {
	var broker = "bc8db60d190e4a3ab672b67319f3c0e9.s1.eu.hivemq.cloud" // find the host name in the Overview of your cluster (see readme)
	var port = 8883                                                    // find the port right under the host name, standard is 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("go-kafka-mqtt") // set a name as you desire
	opts.SetUsername("admin")         // these are the credentials that you declare for your cluster
	opts.SetPassword("Admin123")
	// (optionally) configure callback handlers that get called on certain events
	s.Wg.Add(1)
	opts.OnConnect = s.connectHandler
	opts.OnConnectionLost = connectLostHandler
	// create the client using the options above
	client := mqtt.NewClient(opts)
	// throw an error if the connection isn't successfull
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
