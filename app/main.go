package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
}

func JokeHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	resp, err := http.Get("https://v2.jokeapi.dev/joke/Programming,Pun?blacklistFlags=nsfw&type=single")
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	joke := Joke{}
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(joke.Joke))
}

func main() {
	http.HandleFunc("/", JokeHandler)

	fmt.Println("Starting service on http://127.0.0.1:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
