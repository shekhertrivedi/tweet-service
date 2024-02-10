package main

import (
	"log"

	"github.com/shekhertrivedi/tweet-service/transport"
)

func main() {
	log.Println("Initializing tweet service..!!!")
	transport.InitApp()
}
