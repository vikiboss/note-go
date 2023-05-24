package main

import (
	f "fmt"
	"time"
)

func main() {
	formatStr := "2006-01-02 15:04:05"
	f.Print("Hello Go! And now is ", time.Now().Format(formatStr))
}
