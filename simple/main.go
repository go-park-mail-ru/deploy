package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func MagicNumberHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	magicNumber := rand.Int()
	version := os.Getenv("VERSION")

	w.Header().Set("X-Version", version)

	w.Write([]byte(fmt.Sprintf("Random numer: %d, version: %s", magicNumber, version)))
}

func main() {
	http.HandleFunc("/", MagicNumberHandler)

	fmt.Println("Starting service on http://127.0.0.1:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
