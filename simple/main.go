package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func MagicNumberHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	magicNumber := rand.Int()

	w.Header().Set("X-Version", os.Getenv("VERSION"))

	w.Write([]byte("Magic numer: " + strconv.Itoa(magicNumber)))
}

func main() {
	http.HandleFunc("/", MagicNumberHandler)

	fmt.Println("Starting service on http://127.0.0.1:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
