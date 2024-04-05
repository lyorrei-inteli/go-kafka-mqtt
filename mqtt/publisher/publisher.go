package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Publish(client mqtt.Client) {
	fmt.Println("Publishing MQTT messages")
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Published MQTT message %d", i)
		fmt.Println(text)
		token := client.Publish("movies", 0, false, text)
		token.Wait()
		// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
		if token.Error() != nil {
			fmt.Printf("Failed to publish to topic")
			panic(token.Error())
		}
		time.Sleep(time.Second)
	}
}
