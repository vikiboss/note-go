package main

import (
	"fmt"
)

func main() {
	obj := &struct {
		name string
		age  int
	}{
		name: "Go",
		age:  10,
	}

	myMap := map[string]string{"Go": "go", "Node.js": "node"}

	fmt.Printf("Hello %s, and you are %d years old!\n", obj.name, obj.age)
	fmt.Printf("I love %s and %s\n", myMap["Go"], myMap["Node.js"])
}
