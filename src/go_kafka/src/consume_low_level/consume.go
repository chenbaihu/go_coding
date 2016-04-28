//consumer receive过程（非group）
//1，与任一broker连接
//2，通过topic+partition确定leading broker
//3，从leading broker接收数据

package main

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
)

var topic string
var broker string
var offset int

func init() {
	flag.StringVar(&topic, "topic", "rt_mc", "topic name")
	flag.StringVar(&broker, "broker", "127.0.0.1:9501,127.0.0.1:9501", "broker list")
	flag.IntVar(&offset, "offset", 0, "begin offset")
}

func parseCmdline() error {
	flag.Parse()
	if topic == "" {
		return fmt.Errorf("miss topic name")
	}

	if broker == "" {
		return fmt.Errorf("miss broker list")
	}

	fmt.Printf("topic: %s\n", topic)
	fmt.Printf("broker: %s\n", broker)
	fmt.Printf("offset: %d\n", offset)
	return nil
}

func main() {
	if err := parseCmdline(); err != nil {
		fmt.Printf("%s\n--------------------\n", err)
		flag.PrintDefaults()
		return
	}

	brokerServers := strings.Split(broker, ",")
	fmt.Printf("brokers: %v\n", brokerServers)

	client, err := sarama.NewClient(brokerServers, nil)
	if err != nil {
		fmt.Printf("failed to connect kafka broker, broker=[%s] err=[%s]\n", broker, err)
		return
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		fmt.Printf("failed to create consumer, err=[%s]\n", err)
		return
	}
	defer consumer.Close()

	var partitionID int32 = 0
	partitionConsumer, err := consumer.ConsumePartition(topic, partitionID, int64(offset))
	if err != nil {
		fmt.Printf("failed to consume partition, topic=[%v] partition=[%v] offset=[%v] err=[%v]\n",
			topic, partitionID, offset, err)
		return
	}
	defer partitionConsumer.Close()

	count := 0
	receivedBytes := 0

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			count++
			receivedBytes += len(msg.Value)
			fmt.Printf("receive msg, msg=[%s] partition=[%v] offset=[%v] key=[%s] count=[%v] total_recv_bytes=[%v]\n",
				msg.Value, msg.Partition, msg.Offset, msg.Key, count, receivedBytes)
		case err := <-partitionConsumer.Errors():
			fmt.Printf("receive error, err=[%v]\n", err)
		}
	}
}
