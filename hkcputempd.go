package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/ssimunic/gosensors"
)

func main() {
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		fmt.Println("Error reading sensors, is 'lm-sensors' installed?")
	}

	chipFlag := flag.String("chip", "", "chip name, e.g., thinkpad-isa-0000")
	sensorFlag := flag.String("sensor", "", "sensor name, e.g, temp1")
	pinFlag := flag.String("pin", "32191123", "HomeKit 8-digit PIN for this accessory")
	flag.Parse()

	if *chipFlag == "" || *sensorFlag == "" {
		fmt.Println("USAGE:   " + os.Args[0] + " -pin <accessory PIN> -chip <chip name> -sensor <sensor>")
		fmt.Println("Example: " + os.Args[0] + " -pin 32191123 -chip thinkpad-isa-0000 -sensor temp1")

		fmt.Println("Available Sensors:")
		fmt.Println(sensors.JSON())

		panic("Not enough arguments")
	}

	if _, err := strconv.ParseUint(*pinFlag, 10, 64); err != nil {
		fmt.Println("PIN must be a number")
	}

	_, ok := sensors.Chips[*chipFlag][*sensorFlag]
	if !ok {
		fmt.Println("Available Sensors:")
		fmt.Println(sensors.JSON())

		panic("Chip or sensor not found")
	}

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

			temp_str, ok := sensors.Chips[*chipFlag][*sensorFlag]
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

	t, err := hc.NewIPTransport(hc.Config{Pin: *pinFlag}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
