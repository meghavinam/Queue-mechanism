package main

import (
	"encoding/json"
	"fmt"
	"github.com/remeh/sizedwaitgroup"
	cg "prod/src/customlibrary/configuration"
	er "prod/src/customlibrary/errorhandler"
	sr "prod/src/customlibrary/services"
)

func main() {
	cg.SetConfigParams()

	sr.SetKafkaClient()

	consumer := sr.GetkafkaConsumer()
	sr.ConsumeKafkaMessage(consumer, cg.Config.Kafka.KafkaTopic, deQueueMessage, cg.Config.Kafka.RoutineCount)

}

func deQueueMessage(swg *sizedwaitgroup.SizedWaitGroup, message string) {
	defer swg.Done()
	messageData := make(map[string]interface{})
	err := json.Unmarshal([]byte(message), &messageData)
	fmt.Println(messageData)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		er.ErrorCheck(err)
	}
}
