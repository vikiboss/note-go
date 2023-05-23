package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// output a number between 0 and 9 (not include 10)
	fmt.Print("show you a random number: ", rand.Intn(10))
}
