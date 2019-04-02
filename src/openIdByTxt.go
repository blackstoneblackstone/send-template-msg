package main

import (
	"log"
	"os"
	"time"
	"wxApi"
)

func main() {
	appId := os.Args[1]
	txtPath := os.Args[2]
	wxApi.GetOpenIdFromText(appId, txtPath)
	log.Println("program have finish .....")
	i := 5
	for i < 6 && i > 0 {
		i--
		log.Printf("%d .....", i)
		time.Sleep(time.Second)
	}
}
