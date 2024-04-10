package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	common "prova-sub/internal/common"
)

var messageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	data := common.Payload{}
	payload := msg.Payload()

	if err := json.Unmarshal(payload, &data); err != nil {
		fmt.Println(err)
		return
	}

	if(data.Type == "flores") {
		if data.Humidity > 80 {
			fmt.Printf("%s : %s | %v%%  Umidade [ALERTA: Umidade ALTA] \n", data.EstufaID, data.SensorID, data.Humidity)
		}
		if data.Humidity < 40 {
			fmt.Printf("%s : %s | %v%%  Umidade [ALERTA: Umidade BAIXA] \n", data.EstufaID, data.SensorID, data.Humidity)
		}
		if data.Humidity > 40 && data.Humidity < 80  {
			fmt.Printf("%s : %s | %v%%  Umidade \n", data.EstufaID, data.SensorID, data.Humidity)
		}
	}

	if(data.Type == "hortalicas") {
		if data.Humidity > 70 {
			fmt.Printf("%s : %s | %v%%  Umidade [ALERTA: Umidade ALTA] \n", data.EstufaID, data.SensorID, data.Humidity)
		}
		if data.Humidity < 30 {
			fmt.Printf("%s : %s | %v%%  Umidade [ALERTA: Umidade BAIXA] \n", data.EstufaID, data.SensorID, data.Humidity)
		}
		if data.Humidity > 30 && data.Humidity < 70  {
			fmt.Printf("%s : %s | %v%%  Umidade \n", data.EstufaID, data.SensorID, data.Humidity)
		}
	}
	
}

func main() {

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("subscriber")
	opts.SetDefaultPublishHandler(messageHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("/estufa/#", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
	
	select {}
}