package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func MagicNumberHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	magicNumber := rand.Int()

	w.Write([]byte(strconv.Itoa(magicNumber)))
}

func main() {
	http.HandleFunc("/", MagicNumberHandler)

	fmt.Println("Starting service")

	http.ListenAndServe(":8080", nil)
}
