package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	arena "github.com/xanderseren/arena-go"
)

func getChannel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("ARENA_TOKEN")
	arena := arena.NewClient(token)
	channel, err := arena.GetChannel(p.ByName("channel"), nil)

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

func getBlock(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("ARENA_TOKEN")
	fmt.Println(token)
	arena := arena.NewClient(token)
	block, err := arena.GetBlock(p.ByName("block"), nil)

	if err != nil {
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", err)
	} else {
		ej, _ := json.Marshal(block)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", ej)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/channels/:channel", getChannel)
	router.GET("/blocks/:block", getBlock)
	log.Fatal(http.ListenAndServe(":8080", router))
}
