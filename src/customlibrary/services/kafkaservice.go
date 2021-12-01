package services

import (
	"fmt"
	"github.com/remeh/sizedwaitgroup"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
	cg "prod/src/customlibrary/configuration"
	er "prod/src/customlibrary/errorhandler"
)

var KafkaProducer *kafka.Producer
var KafkaConsumer *kafka.Consumer

func SetkafkaProducer() {
	broker := cg.Config.Kafka.KafkaBrokerIp
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	er.ErrorCheck(err)
	KafkaProducer = p
}

func GetkafkaProducer() *kafka.Producer {
	return KafkaProducer
}
func AddToKafka(p *kafka.Producer, topic string, value string) {

	deliveryChan := make(chan kafka.Event)
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	}, deliveryChan)

	if err != nil {
		fmt.Println(err)
		close(deliveryChan)
		er.ErrorCheck(err)
		return
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		// fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
		// 	*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(deliveryChan)

}

func SetKafkaClient() {
	if cg.Config.Kafka.KafkaBrokerIp == "" {
		fmt.Println("Kafka configuration not set")
		os.Exit(1)
	}
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cg.Config.Kafka.KafkaBrokerIp,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	KafkaConsumer = c
}

func GetkafkaConsumer() *kafka.Consumer {
	return KafkaConsumer
}

func ConsumeKafkaMessage(c *kafka.Consumer, topic string, handler func(swg *sizedwaitgroup.SizedWaitGroup, message string), concurrency int) {
	c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)
	swg := sizedwaitgroup.New(concurrency)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			swg.Add()
			go handler(&swg, string(msg.Value))
		} else {
			// fmt.Printf("Consumer error: %v (%v)\n", err, msg, topic)
			// os.Exit(1)
			// The client will automatically try to recover from all errors.

		}
	}
	swg.Wait()
}
