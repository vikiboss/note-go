package main

import (
	"fmt"
)

func main() {
	sayHi()
	sayHello := saySomeFactory("Hello ", "!")
	sayHello("Go")
}

func sayHi() {
	fmt.Println("Hi Go!")
}

func saySomeFactory(str, end string) func(content string) {
	return func(content string) {
		fmt.Println(str + content + end)
	}
}
