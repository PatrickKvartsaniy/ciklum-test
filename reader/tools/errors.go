package tools

import (
	"log"
)

// CheckErr - function for  errors handling
func CheckErr(err error) {
	if err != nil {
		// panic(err)
		log.Fatal(err.Error())

	}
}
