package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func OilPriceHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	magicNumber := rand.Int()

	w.Write([]byte("Oil price 2.0: " + strconv.Itoa(magicNumber)))
}

func main() {
	http.HandleFunc("/", OilPriceHandler)

	fmt.Println("Starting service")
	http.ListenAndServe(":8080", nil)
}
