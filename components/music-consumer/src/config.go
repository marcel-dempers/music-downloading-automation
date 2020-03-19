package main

import (
	"github.com/music-consumer/models"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

func GetConfiguration() *models.Configuration{

	file, e := ioutil.ReadFile("/app/configs/config.json")
    if e != nil {
        fmt.Printf("Read config file error: %v\n", e)
        os.Exit(1)
	}

	var config *models.Configuration
	err := json.Unmarshal(file, &config)
	
	if err != nil {
        fmt.Printf("Problem in config file: %v\n", err.Error())
        os.Exit(1)
	}

	return config
}
