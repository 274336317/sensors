package main 

import (
"log"
	"time"
)

func main() {
	log.Print("Sensor Start Running...")
	initSensors()
	go RunSensors()
	go StartMqttClient()
	for ;;{
		time.Sleep(time.Duration(3000)*time.Millisecond)
	}	
}

