package common

import (
	"fmt"
	"math/rand"
	"time"
	"math"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)


type Payload struct {
	EstufaID          string `json:"estufa_id"`
	SensorID		  string `json:"sensor_id"`
	Type        string `json:"type"`
	Humidity 		float64 `json:"humidity"`
	TimeStamp   string `json:"timestamp"`
}

type MaxMin struct {
	Minimum float64
	Maximum float64
}


func PayloadValueGenerator() float64 {
	rand.NewSource(time.Now().UnixNano())
	max := 100.0
	min := 0.0
	value := rand.Float64()*(max-min) + min
	return math.Round(value)
}

func ConnectMQTT(seed rand.Source) MQTT.Client {

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID(fmt.Sprintf("greenhouse-%d", seed.Int63()))
	client := MQTT.NewClient(opts)
	if session := client.Connect(); session.Wait() && session.Error() != nil {
		panic(session.Error())
	}
	
	return client
}

