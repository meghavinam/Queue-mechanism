package main

import (
	"encoding/json"
	cg "prod/src/customlibrary/configuration"
	er "prod/src/customlibrary/errorhandler"
	sr "prod/src/customlibrary/services"
	"sync"
)

func main() {
	cg.SetConfigParams()
	sr.SetkafkaProducer()
	var wg sync.WaitGroup
	messageArray := map[string]interface{}{}
	for i := 0; i <= cg.Config.Kafka.MessageLength; i++ {
		messageArray["message"] = "hii"
		messageArray["counter"] = i
		wg.Add(1)
		go Enqueue(messageArray, &wg)
		wg.Wait()
	}

}

func Enqueue(messageArray map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	producer := sr.GetkafkaProducer()
	jsonData, err := json.Marshal(messageArray)
	checkErr(err)
	sr.AddToKafka(producer, cg.Config.Kafka.KafkaTopic, string(jsonData))

}

func checkErr(err error) {
	if err != nil {
		er.ErrorCheck(err)
	}
}
