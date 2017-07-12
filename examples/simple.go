package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	arena "github.com/xanderseren/arena-go"
)

func getChannel(w http.ResponseWriter, r *http.Request) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	arena := arena.NewClient()
	channel, err := arena.GetChannel("60260")

	if err != nil {
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", err)
	} else {
		ej, _ := json.Marshal(channel)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", ej)
	}
}

func main() {
	http.HandleFunc("/", getChannel)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
