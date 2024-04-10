package greenhouse

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	common "prova-sub/internal/common"
)


func Flores(estufa_id int, sensor_id int) {

	client := common.ConnectMQTT(rand.New(rand.NewSource(time.Now().UnixNano())))

	for {
		data := common.Payload{
			EstufaID:          fmt.Sprintf("estufa-%v", estufa_id),
			SensorID: 			fmt.Sprintf("flores-%v", sensor_id),
			Type:        "flores",
			Humidity:       common.PayloadValueGenerator(),
			TimeStamp:   time.Now().String(),
		}

		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return
		}

		token := client.Publish("/estufa/flores", 1, false, string(payload))
		fmt.Println("Publishing message to /flores")
		token.Wait()
		time.Sleep(3 * time.Second)
	}
}

func Hortalicas(estufa_id int, sensor_id int) {

	client := common.ConnectMQTT(rand.New(rand.NewSource(time.Now().UnixNano())))

	for {
		data := common.Payload{
			EstufaID:     fmt.Sprintf("estufa-%v", estufa_id),
			SensorID:     fmt.Sprintf("hortalicas-%v", sensor_id),
			Type:        "hortalicas",
			Humidity:    common.PayloadValueGenerator(),
			TimeStamp:   time.Now().String(),
		}

		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return
		}

		token := client.Publish("/estufa/hortalicas", 1, false, string(payload))
		fmt.Println("Publishing message to /hortalicas")
		token.Wait()
		time.Sleep(3 * time.Second)
	}
}
