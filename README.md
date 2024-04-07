# Ponte Kafka-MQTT em Go

Essa aplicação em Go proporciona a integração entre o cluster do HiveMQ e um provedor Kafka na nuvem, estabelecendo uma ponte eficiente para o fluxo de mensagens entre os dois serviços.

## Índice
- [Pré-requisitos](#pré-requisitos)
- [Iniciando](#iniciando)
- [Configuração](#configuração)
- [Executando a Aplicação](#executando-a-aplicação)
- [Testes](#testes)
- [Demonstração](#demonstração)

## Pré-requisitos
- Instância Kafka em execução (HiveMQ ou outro provedor na nuvem)
- Instância MQTT broker em execução
- Go instalado na máquina local

## Iniciando
[Repositório GitHub](https://github.com/lyorrei-inteli/go-kafka-mqtt)

Clone o repositório para sua máquina local.

## Configuração
Configure seu broker Kafka e MQTT no arquivo `client.properties` e `mqtt/config/config.go` respectivamente.

client.properties:
```
bootstrap.servers=...
security.protocol=...
sasl.mechanisms=...
sasl.username=...
sasl.password=...
session.timeout.ms=...
```

mqtt/config/config.go:
```go
var broker = "..."
var port = 8883
opts := mqtt.NewClientOptions()
opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
opts.SetClientID("...")
opts.SetUsername("...")
opts.SetPassword("...")
```

## Executando a Aplicação
Inicie a ponte com o comando:

```bash
go run .
```

## Testes
Para validar a integração, execute:

```bash
go test ./tests -v
```

## Demonstração
Para visualizar o funcionamento do sistema, assista ao vídeo demonstrativo [aqui](https://youtu.be/OZuMXg120Iw).<br/>
ou <br/>
https://youtu.be/OZuMXg120Iw