package main

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"log"
	"time"
)

func main() {
	info := model.Info{
		Name: "My Door Sensor",
	    SerialNumber: "XXX-YYY-ZZZ",
	    Manufacturer: "Me",
	    Model: "1",
	    Firmware: "1.0",
	}

	doorSensor := accessory.NewContactSensor(info)

	t, err := hap.NewIPTransport(hap.Config{Pin: "32191123"}, doorSensor.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	// Periodically toggle the switch's on characteristic
	go func() {
		for {
			var state model.ContactSensorStateType
			if doorSensor.State() == model.ContactDetected {
				state = model.ContactNotDetected
			} else {
				state = model.ContactDetected
			}
			doorSensor.SetState(state)
			time.Sleep(5 * time.Second)
		}
	}()

	hap.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}