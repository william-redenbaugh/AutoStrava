package main

import (
	"log"

	"github.com/3nueves/serial"
)

/*
	@brief Setting up serial device, flags if there is an error
*/
func open_serial_port(com_port string) *serial.Port {
	c := &serial.Config{Name: com_port, Baud: 115200}

	s, err := serial.OpenPort(c)
	check_err(err)
	return s
}

/*
	@brief Sets up communication with the embedded device.
*/
func setup_communication() {

	s := open_serial_port("/dev/tty/USB0")

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}
