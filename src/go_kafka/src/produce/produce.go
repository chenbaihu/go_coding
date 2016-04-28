package main

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"time"
)

var topic string
var broker string
var count int
var data []byte = []byte("kafka_high_level_produce_data_testssssssssssssssssssssss")

func init() {
	flag.StringVar(&topic, "topic", "rt_mc", "topic name")
	flag.StringVar(&broker, "broker", "127.0.0.1:9501,127.0.0.1:9501", "broker list")
	flag.IntVar(&count, "count", 1, "count of messages")
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
	fmt.Printf("count of messages: %d\n", count)
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

	producer, err := sarama.NewSyncProducerFromClient(client)
	//config := sarama.NewConfig()
	//config.Producer.Return.Errors = true
	//config.Producer.Return.Successes = true
	//config.Producer.Flush.Messages = *batchNum
	//config.Producer.Flush.MaxMessages = *batchNum
	//if *partition != -1 {
	//	config.Producer.Partitioner = sarama.NewManualPartitioner
	//}
	////config.Producer.Compression = 2
	//if *waitForAll {
	//	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	//} else {
	//	config.Producer.RequiredAcks = sarama.WaitForLocal
	//}
	//config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	//producer, err := sarama.NewSyncProducer(brokerServers, config)
	if err != nil {
		fmt.Printf("failed to create sync producer, err=[%s]\n", err)
		return
	}
	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}

	//指定partitionid的写入
	//msg := &sarama.ProducerMessage{
	//	Topic:     topic,
	//	Key:       sarama.StringEncoder(fmt.Sprintf("%d", start/1000000)),
	//	Value:     sarama.StringEncoder(data),
	//	Metadata:  &MessageMetadata{EnqueuedAt: time.Now()},
	//	Partition: (int32)(*partition),
	//}

	//不指定partitionid系统自动负载均衡
	//msg := &sarama.ProducerMessage{
	//	Topic:    topic,
	//	Key:      sarama.StringEncoder(fmt.Sprintf("%d", start/1000000)),
	//	Value:    sarama.StringEncoder(data),
	//	Metadata: &MessageMetadata{EnqueuedAt: time.Now()},
	//}

	sendCount := 0
	min := 1000.0
	max := 0.0

	start := time.Now().UnixNano()
	for {
		if count >= 0 && sendCount >= count {
			break
		}
		tmp := time.Now().UnixNano()
		partition, offset, err := producer.SendMessage(&msg)
		elapsed := float64(time.Now().UnixNano()-tmp) / 1.0e6 //ms
		if elapsed < min {
			min = elapsed
		}

		if elapsed > max {
			max = elapsed
		}

		if err != nil {
			fmt.Printf("send failed, err=[%s]\n", err)
			return
		}

		fmt.Printf("send succeed, partition=%v offset=%v msg=[%v]\n", partition, offset, msg)
		sendCount++
	}
	elapsed := float64(time.Now().UnixNano()-start) / 1.0e6 //ms
	sendBytes := len(data) * count                          // B
	rateByte := float64(sendBytes) / elapsed                //KB/s
	rateMsg := float64(count) / elapsed                     //K/s
	avg := elapsed / float64(count)                         //ms

	fmt.Printf("count=[%v] time=[%vms] bytes=[%v] rate=[%vKB/s(%vK/s)] max=[%vms] min=[%vms] avg=[%vms]\n",
		count, elapsed, sendBytes, rateByte, rateMsg, max, min, avg)
}
