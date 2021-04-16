package main

import (
	"github.com/3nueves/serial"
)

/*!
@brief LED matrix struct that let's us control stuff
*/
type RGBColor struct {
	red   uint8
	green uint8
	blue  uint8
}

/*!
@brief Data structure that helps manipulate LED matrix manipulation
*/
type LEDMatrixSerial struct {
	serial_device *serial.Port
	output_color  [16][8]RGBColor
}

/*!
@brief Easy manipulation of pixel colors
*/
func (serial LEDMatrixSerial) setPixelColor(x uint8, y uint8, r uint8, g uint8, b uint8) {
	serial.output_color[x][y].red = r
	serial.output_color[x][y].green = g
	serial.output_color[x][y].blue = b
}

/*!
@brief pushes up LED matrix data to serial device.
*/
func (serial LEDMatrixSerial) update() {
	output_signature := []byte{20, 30, 40, 50}
	serial.serial_device.Write(output_signature)

	for x := 0; x < 15; x++ {
		for y := 0; y < 8; y++ {
			pixel_array := []byte{serial.output_color[x][y].red, serial.output_color[x][y].green, serial.output_color[x][y].blue}
			serial.serial_device.Write(pixel_array)
		}
	}
}

/*!
@brief Setting up serial device, flags if there is an error
*/
func open_serial_port(com_port string) *serial.Port {
	c := &serial.Config{Name: com_port, Baud: 115200}

	s, err := serial.OpenPort(c)
	check_err(err)
	return s
}

/*!
@brief Sets up communication with the embedded device.
*/
func setup_communication(port string) LEDMatrixSerial {
	var led_matrix LEDMatrixSerial
	s := open_serial_port(port)
	led_matrix.serial_device = s
	return led_matrix
}
