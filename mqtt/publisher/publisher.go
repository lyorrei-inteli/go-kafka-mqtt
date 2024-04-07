package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Publish(client mqtt.Client) {
	fmt.Println("Publishing MQTT messages")

	// Array de filmes para simular a publicação alternada
	movies := []string{"The Shawshank Redemption", "The Godfather", "The Dark Knight", "12 Angry Men", "Schindler's List"}

	// Número de mensagens a serem publicadas
	num := len(movies)

	for i := 0; i < num; i++ {
		// Seleciona um filme do array com base no índice atual
		movie := movies[i]

		// Formata a mensagem para incluir o nome do filme
		text := fmt.Sprintf("Published movie in MQTT: %s", movie)

		fmt.Println(text)

		// Publica a mensagem no tópico "movies"
		token := client.Publish("movies", 0, false, text)
		token.Wait()

		// Verifica se houve erros durante a publicação
		if token.Error() != nil {
			fmt.Printf("Failed to publish to topic")
			panic(token.Error())
		}

		// Pausa antes da próxima publicação
		time.Sleep(time.Second)
	}
}
