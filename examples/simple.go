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

func getBlock(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("ARENA_TOKEN")
	arena := arena.NewClient(token)
	block, err := arena.Blocks.Get(p.ByName("block"), nil)
	fmt.Println(block.Title)

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

func postBlock(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grabs this channel https://api.are.na/v2/channels/golang
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("ARENA_TOKEN")
	a := arena.NewClient(token)
	block, err := a.Blocks.AddSource(p.ByName("channel"), "http://youtube.com")

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
	router.GET("/blocks/:block", getBlock)
	router.GET("/search/blocks", searchBlock)
	router.POST("/channels/:channel/blocks", postBlock)
	// router.GET("/search", search)
	// router.PUT("/blocks/:block", editBlock)
	log.Fatal(http.ListenAndServe(":8080", router))
}
