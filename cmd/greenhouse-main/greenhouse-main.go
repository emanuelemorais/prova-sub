package main

import (
	greenhouse "prova-sub/internal/greenhouse"
	"sync"
)


func main() {

	numEstufas := 4
	numSensoresFlores := 3
	numSendoresHortalicas := 3
	var wgFlores sync.WaitGroup
	var wgHortalicas sync.WaitGroup


	for i := 1; i < numEstufas; i++ {
		for j := 0; j < numSensoresFlores; j++ {
			wgFlores.Add(1)
			go func(estufa_id int, sensor_id int) {
				defer wgFlores.Done()
				greenhouse.Flores(estufa_id, sensor_id)
			}(i, j)
		}
		for k := 0; k < numSendoresHortalicas; k++ {
			wgHortalicas.Add(1)
			go func(estufa_id int, sensor_id int) {
				defer wgHortalicas.Done()
				greenhouse.Hortalicas(estufa_id, sensor_id)
			}(i, k)
		}
	}

	wgFlores.Wait()
	wgHortalicas.Wait()
}