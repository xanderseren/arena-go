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

type EditBlockStruct struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Content     *string `json:"content"`
}

func getChannel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("ARENA_TOKEN")
	arena := arena.NewClient(token)
	channel, err := arena.Channels.Get(p.ByName("channel"), nil)

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

func getChannelContents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("ARENA_TOKEN")
	arena := arena.NewClient(token)
	channel, err := arena.Channels.Contents(p.ByName("channel"), nil)

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

func editTitle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	token := os.Getenv("ARENA_TOKEN")
	a := arena.NewClient(token)
	err := a.Blocks.EditTitle(p.ByName("block"), "New Title")

	if err != nil {
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(204)
	}
}

func searchBlock(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	queryValues := r.URL.Query()
	q := queryValues.Get("q")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("ARENA_TOKEN")
	a := arena.NewClient(token)
	block, err := a.Blocks.Search(arena.Arguments{"q": q})

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

func addChannel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("ARENA_TOKEN")
	a := arena.NewClient(token)
	block, err := a.Channels.Add("Testing new channel", "private")

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

// func search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	// Grabs this channel https://api.are.na/v2/channels/golang
// 	queryValues := r.URL.Query()
// 	q := queryValues.Get("q")
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
//
// 	token := os.Getenv("ARENA_TOKEN")
// 	a := arena.NewClient(token)
// 	block, err := a.Search(arena.Arguments{"q": q})
//
// 	if err != nil {
// 		w.WriteHeader(200)
// 		fmt.Fprintf(w, "%s", err)
// 	} else {
// 		ej, _ := json.Marshal(block)
//
// 		// Write content-type, statuscode, payload
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(200)
// 		fmt.Fprintf(w, "%s", ej)
// 	}
// }

// func editBlock(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	// Grabs this channel https://api.are.na/v2/channels/golang
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
//
// 	e := EditBlockStruct{}
//
// 	token := os.Getenv("ARENA_TOKEN")
// 	arena := arena.NewClient(token)
// 	block, err := arena.EditBlock(p.ByName("block"), nil)
//
// 	if err != nil {
//     w.WriteHeader(200)
// 		fmt.Fprintf(w, "%s", err)
// 	} else {
// 		ej, _ := json.Marshal(block)
//
// 		// Write content-type, statuscode, payload
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(200)
// 		fmt.Fprintf(w, "%s", ej)
// 	}
// }

func main() {
	router := httprouter.New()
	router.GET("/channels/:channel", getChannel)
	router.GET("/channels/:channel/contents", getChannelContents)
	router.PUT("/blocks/:block", editTitle)
	router.GET("/search/blocks", searchBlock)
	router.POST("/channels/", addChannel)
	// router.GET("/search", search)
	// router.PUT("/blocks/:block", editBlock)
	log.Fatal(http.ListenAndServe(":8080", router))
}
