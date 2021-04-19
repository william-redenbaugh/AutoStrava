package main

import "time"

func main() {
	//get_weather_data()

	matrix := setup_communication("/dev/ttyACM0")

	for {
		matrix.drawChar(uint8(0), uint8(0), uint8(100), uint8(100), uint8(100), uint8('0'))
		matrix.update()

		time.Sleep(time.Second)
	}
}
