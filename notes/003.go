package main

import (
	"fmt"
	"math/rand"
)

const tip string = "show you a random number: "

func main() {
	// output a number between 0 and 9 (not include 10)
	fmt.Print(tip, rand.Intn(10))
}
