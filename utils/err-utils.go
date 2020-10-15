package utils

import "log"

// Error check utils ...

func PanicErr(err error) {
	if err != nil {
		log.Println("Panic ::")
		panic(err)
	}
}

func FatalErr(err error) {
	if err != nil {
		log.Println("Fatal ::")
		log.Fatal(err)
	}
}

func LogErr(err error) {
	if err != nil {
		log.Println("An error occurred :: ", err)
	}
}
