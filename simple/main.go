package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	magicNumber := rand.Int()

	if magicNumber == 5 {
		fmt.Println("Oil price is negative =(")
		return
	}

	fmt.Println("=|")
}
