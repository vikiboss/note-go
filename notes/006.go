package main

import (
	"fmt"
)

func main() {
	obj := &struct {
		name      string
		birthYear int
	}{
		name:      "Go",
		birthYear: 2009,
	}

	myMap := map[string]string{"Go": "go", "Node.js": "node"}

	fmt.Printf("Hello %s! Your birth year is %d!\n", obj.name, obj.birthYear)
	fmt.Printf("I love %s and %s\n", myMap["Go"], myMap["Node.js"])
}
