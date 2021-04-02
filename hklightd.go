package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/ssimunic/gosensors"
)

func main() {
	sensors, err := gosensors.NewFromSystem()
	fmt.Println("Available Sensors:")
	fmt.Println(sensors.JSON())

	info := accessory.Info{
		Name:         "Server Temperature",
		Manufacturer: "custom",
	}

	temp := 1.0
	acc := accessory.NewTemperatureSensor(info, temp, -50.0, +200.0, 0.1)
	go func() {
		for {
			// this calls `sensors` and returns the output
			sensors, err := gosensors.NewFromSystem()
			if err != nil {
				panic(err)
			}

			temp_str, ok := sensors.Chips["thinkpad-isa-0000"]["temp1"]
			if !ok {
				panic("Sensor not found")
			}

			var temp float64
			fmt.Sscanf(temp_str, "%f", &temp)
			fmt.Println(temp)
			if err == nil {
				acc.TempSensor.CurrentTemperature.SetValue(temp)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	t, err := hc.NewIPTransport(hc.Config{Pin: "32191123"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
