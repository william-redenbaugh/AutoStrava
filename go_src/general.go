package main

import (
	"log"
)

/*
	@brief Check and flags errors.
*/
func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
