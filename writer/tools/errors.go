package tools

import "log"

func CheckErr(err error) {
	if err != nil {
		// panic(err)
		log.Println(err.Error())
	}
}
