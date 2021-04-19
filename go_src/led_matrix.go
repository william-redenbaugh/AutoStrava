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
func (serial *LEDMatrixSerial) setPixelColor(x uint8, y uint8, r uint8, g uint8, b uint8) {

	// Out of bounds system
	if x > 15 || y > 7 {
		return
	}

	serial.output_color[x][y].red = r
	serial.output_color[x][y].green = g
	serial.output_color[x][y].blue = b
}

func (serial *LEDMatrixSerial) drawChar(x uint8, y uint8, r uint8, g uint8, b uint8, characterCh uint8) {

	switch characterCh {
	case '0':
		// Drawing top piece
		serial.setPixelColor(1+x, 0+y, r, g, b)
		serial.setPixelColor(2+x, 0+y, r, g, b)
		serial.setPixelColor(3+x, 0+y, r, g, b)

		// Drawing the left side
		serial.setPixelColor(x, 1+y, r, g, b)
		serial.setPixelColor(x, 2+y, r, g, b)
		serial.setPixelColor(x, 3+y, r, g, b)
		serial.setPixelColor(x, 4+y, r, g, b)

		// Drawing the bottom side
		serial.setPixelColor(1+x, 5+y, r, g, b)
		serial.setPixelColor(2+x, 5+y, r, g, b)
		serial.setPixelColor(3+x, 5+y, r, g, b)

		// Drawing the right side
		serial.setPixelColor(4+x, 1+y, r, g, b)
		serial.setPixelColor(4+x, 2+y, r, g, b)
		serial.setPixelColor(4+x, 3+y, r, g, b)
		serial.setPixelColor(4+x, 4+y, r, g, b)

	case '1':

	case '2':
	case '3':
	case '4':
	case '5':
	case '6':
	case '7':
	case '8':
	case '9':

	}

}

/*!
@brief pushes up LED matrix data to serial device.
*/
func (serial LEDMatrixSerial) update() {

	pixel_array := []byte{16, 24, 33, 22}
	serial.serial_device.Write(pixel_array)

	for x := 0; x < 16; x++ {
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
