package main

import (
	"log"
	"twitter_gin/src/api/app"
)

func main() {
	if err := app.StartApp(); err != nil {
		log.Fatal("error starting server", err)
	}
}
