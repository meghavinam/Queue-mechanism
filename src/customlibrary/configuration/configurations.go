package configuration

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Configuration struct {
	Kafka struct {
		KafkaBrokerIp string
		RoutineCount  int
		KafkaTopic    string
		MessageLength int
	}

	Jira struct {
		Host       string
		Username   string
		ApiKey     string
		ProjectKey string
		AssigneeId string
	}
}

var Config = Configuration{}

func SetConfigParams() {
	_, b, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(b)
	c := flag.String("c", currentPath+"/../../config/config.json", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
}
