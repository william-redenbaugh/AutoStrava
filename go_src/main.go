package main

import "time"

func main() {
	//get_weather_data()

	matrix := setup_communication("/dev/ttyACM0")

	for x := 0; x < 16; x++ {
		for y := 0; y < 8; y++ {
			matrix.setPixelColor(uint8(x), uint8(y), 100, 100, 100)
			matrix.update()

			time.Sleep(time.Millisecond * 10)
		}
	}

	for x := 0; x < 16; x++ {
		for y := 0; y < 8; y++ {
			matrix.setPixelColor(uint8(x), uint8(y), 0, 0, 0)
			matrix.update()

			time.Sleep(time.Millisecond * 10)
		}
	}
}
