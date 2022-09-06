package main

import (
    // "net/url"
	"encoding/json"
	"net/http"
	"log"
	// "poker/game"
)

func main() {
	// game.Start()
	resp, err := http.Get("http://172.26.0.3:8080")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	var j interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		log.Println(err)
	}
	log.Println(j)
}
