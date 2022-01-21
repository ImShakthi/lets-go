package snippets

import (
	"flag"
	"testing"
)

var broker = flag.String("broker", "127.0.0.1", "broker IP addr")
var topic = flag.String("topic", "", "kafka topic name")
var value = flag.String("data", "", "payload data")
var help = flag.Bool("help", false, "help")
var headerData = flag.String("header", "", "header data")

func TestProduceMsg(t *testing.T) {

	produceMsg(*broker, *topic, *value, *help, *headerData)
}
