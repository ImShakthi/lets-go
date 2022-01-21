package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"os"
)

const (
	helpline = `go run main.go -broker <ip-addr> -topic <topic-name> -data '<sample data>' -header '[{"key":"k1","data":"v1"},{"key":"k2","data":"v2"}]'
-broker, -topic, -data/-file are mandatory flags 
Examples:
go run main.go -broker 127.0.0.1 -topic test-topic-name -data '{"msg":"hello"}'
go run main.go -broker 127.0.0.1 -topic test-topic-name -file '~/Desktop/data.json'`
)

func main() {
	broker := flag.String("broker", "127.0.0.1", "broker IP addr")
	topic := flag.String("topic", "", "kafka topic name")
	data := flag.String("data", "", "payload data")
	help := flag.Bool("help", false, "help")
	headerData := flag.String("header", "", "header data")
	filePath := flag.String("file", "", "file path of payload")
	flag.Parse()

	args := cliArgs{
		broker:     *broker,
		topic:      *topic,
		data:       *data,
		help:       *help,
		headerData: *headerData,
		filePath:   *filePath,
	}

	if args.help {
		fmt.Fprintf(os.Stdout, helpline)
		os.Exit(0)
	}

	err := args.validate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	payload, err := args.getPayload()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	kafkaHeaders, err := args.getHeaders()
	fmt.Fprintf(os.Stdout, "kafkaHeaders= %+v\n", kafkaHeaders)

	configMap := kafka.ConfigMap{
		"bootstrap.servers": args.broker,
		// add config here
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
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
		Value:          []byte(payload),
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

type cliArgs struct {
	broker     string
	topic      string
	data       string
	help       bool
	headerData string
	filePath   string
}

func (args cliArgs) validate() error {
	if len(args.topic) == 0 {
		return fmt.Errorf("topic can't be empty")
	}

	if len(args.data) == 0 && len(args.filePath) == 0 {
		return fmt.Errorf("data and file path can't be empty")
	}

	if len(args.data) != 0 && len(args.filePath) != 0 {
		return fmt.Errorf("data or file path, one only should be used")
	}

	return nil
}

func (args cliArgs) getPayload() (string, error) {
	if len(args.filePath) != 0 {
		fileContent, err := ioutil.ReadFile(args.filePath)
		if err != nil {
			return "", fmt.Errorf("error in reading data from %s due to %+v \n", args.filePath, err)
		}
		return string(fileContent), nil
	}
	return args.data, nil
}

func (args cliArgs) getHeaders() ([]kafka.Header, error) {
	if len(args.headerData) == 0 {
		return []kafka.Header{}, nil
	}

	var headers Headers
	err := json.Unmarshal([]byte(args.headerData), &headers)
	if err != nil {
		return []kafka.Header{}, fmt.Errorf("unable to parse header due to %+v", err)
	}
	return headers.mapTo(), nil
}

type Headers []Header

type Header struct {
	Key   string `json:"key"`
	Value string `json:"data"`
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
