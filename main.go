package main

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
)

func main() {
	for range time.Tick(2 * time.Second) {
		execute()
	}
}

func execute() {
	i2c, err := i2c.NewI2C(0x76, 1)
	if err != nil {
		log.Err(err).Msg("new sensor")
		return
	}
	defer i2c.Close()

	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c)
	if err != nil {
		log.Err(err).Msg("new sensor")
		return
	}
	for range time.Tick(10 * time.Second) {
		v, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_HIGHEST)
		if err != nil {
			log.Err(err).Msg("read temp")
			return
		}
		log.Printf("Temprature = %v*C\n", v)
		_, v, err = sensor.ReadHumidityRH(bsbmp.ACCURACY_HIGHEST)
		if err != nil {
			log.Err(err).Msg("read humidity")
			return
		}
		log.Printf("Humidity = %v RH\n", v)
		v, err = sensor.ReadPressurePa(bsbmp.ACCURACY_HIGHEST)
		if err != nil {
			log.Err(err).Msg("read pressure pa")
			return
		}
		log.Printf("Pressure = %v Pa\n", v)
		v, err = sensor.ReadPressureMmHg(bsbmp.ACCURACY_HIGHEST)
		if err != nil {
			log.Err(err).Msg("read pressure mmhg")
			return
		}
		log.Printf("Pressure = %v mmHg\n", v)
		v, err = sensor.ReadAltitude(bsbmp.ACCURACY_HIGHEST)
		if err != nil {
			log.Err(err).Msg("read altitude")
			return
		}
		log.Printf("Altitude = %v m\n", v)
	}
}
