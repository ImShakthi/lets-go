package snippets

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

/*
go run main.go -broker <ip-addr> -topic <topic-name> -data '<sample data>'

example
go run main.go -broker 127.0.0.1 -topic test-topic-name -data '{"msg":"hello"}'
*/

func produceMessage() {
	broker := flag.String("broker", "127.0.0.1", "broker IP addr")
	topic := flag.String("topic", "", "kafka topic name")
	value := flag.String("data", "", "payload data")
	help := flag.Bool("help", false, "help")
	headerData := flag.String("header", "", "header data")
	flag.Parse()

	produceMsg(*broker, *topic, *value, *help, *headerData)
}

func produceMsg(broker string, topic string, value string, help bool, headerData string) {
	if help {
		helpline := `go run main.go -broker <ip-addr> -topic <topic-name> -data '<sample data>' -header '[{"key":"k1","value":"v1"},{"key":"k2","value":"v2"}]'
-broker, -topic, -data are mandatory flags 
Example: go run main.go -broker 127.0.0.1 -topic test-topic-name -data '{"msg":"hello"}'`

		fmt.Fprintf(os.Stdout, helpline)
		os.Exit(0)
	}

	if len(topic) == 0 {
		fmt.Fprintf(os.Stderr, "topic can't be empty\n")
		os.Exit(1)
	}

	if len(value) == 0 {
		fmt.Fprintf(os.Stderr, "data can't be empty\n")
		os.Exit(1)
	}

	if len(headerData) == 0 {
		fmt.Fprintf(os.Stderr, "data can't be empty\n")
		os.Exit(1)
	}

	var kafkaHeaders []kafka.Header
	if len(headerData) != 0 {
		var headers Headers
		err := json.Unmarshal([]byte(headerData), &headers)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to parse header %+v\n", err)
			os.Exit(1)
		}
		kafkaHeaders = headers.mapTo()
	}
	fmt.Fprintf(os.Stdout, "kafkaHeaders= %+v\n", kafkaHeaders)

	configMap := kafka.ConfigMap{
		"bootstrap.servers": broker,
	}

	p, err := kafka.NewProducer(&configMap)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created Producer %v\n", p)

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	message := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
		Headers:        kafkaHeaders,
	}
	err = p.Produce(&message, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}

type Headers []Header

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (h Headers) mapTo() []kafka.Header {
	headers := make([]kafka.Header, 0)
	for _, header := range h {
		headers = append(headers, kafka.Header{
			Key:   header.Key,
			Value: []byte(header.Value),
		})
	}
	return headers
}
