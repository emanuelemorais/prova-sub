package greenhouse

import (
	"fmt"
	"testing"
	"time"
	"encoding/json"
	common "prova-sub/internal/common"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func validatePayload(payload common.Payload) bool {
	if payload.EstufaID == "" || payload.SensorID == "" || payload.TimeStamp == "" || payload.Humidity == 0 || payload.Type == "" {
		return false
	}
	return true
}

func TestFlores(t *testing.T) {


	var handler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("New message with qos %d on topic %s: %s \n", msg.Qos(), msg.Topic(), msg.Payload())

		data := common.Payload{}
		payload := msg.Payload()

		if err := json.Unmarshal(payload, &data); err != nil {
			t.Errorf("Error unmarshalling")
			fmt.Println(err)
			return
		}

		validatedFields :=validatePayload(data)
		if !validatedFields {  
			t.Errorf("One field or more of the payload is empty")
		}

		if msg.Qos() != 1 {
			t.Errorf("QoS must be 1")
		}
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("test")
	opts.SetDefaultPublishHandler(handler)

	client := MQTT.NewClient(opts)
	defer client.Disconnect(500)

	if session := client.Connect(); session.Wait() && session.Error() != nil {
		panic(session.Error())
	}

	done := make(chan bool)

	go func() {
		Flores(1, 1) //Send with QOS 1
		done <- true
	}()

	go func() {
		if token := client.Subscribe("/estufa/flores", 1, nil); token.Wait() && token.Error() != nil {
			t.Logf("Error subscribing: %s", token.Error())
			return
		}
	}()

	time.Sleep(4 * time.Second)

}

func TestQOSHortalicas(t *testing.T) {

	var handler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("New message with qos %d on topic %s: %s \n", msg.Qos(), msg.Topic(), msg.Payload())

		data := common.Payload{}
		payload := msg.Payload()

		if err := json.Unmarshal(payload, &data); err != nil {
			t.Errorf("Error unmarshalling")
			fmt.Println(err)
			return
		}

		validatedFields :=validatePayload(data)
		if !validatedFields {  
			t.Errorf("One field or more of the payload is empty")
		}

		if msg.Qos() != 1 {
			t.Errorf("QoS must be 1")
		}
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("test")
	opts.SetDefaultPublishHandler(handler)

	client := MQTT.NewClient(opts)
	defer client.Disconnect(500)

	if session := client.Connect(); session.Wait() && session.Error() != nil {
		panic(session.Error())
	}

	done := make(chan bool)

	go func() {
		Hortalicas(1, 1) //Send with QOS 1
		done <- true
	}()

	go func() {
		if token := client.Subscribe("/estufa/hortalicas", 1, nil); token.Wait() && token.Error() != nil {
			t.Logf("Error subscribing: %s", token.Error())
			return
		}
	}()

	time.Sleep(4 * time.Second)

}