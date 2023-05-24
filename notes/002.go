package main

import (
	"fmt"
	"time"
)

func main() {
	formatStr := "2006-01-02 15:04:05"
	fmt.Print("Hello Go! And now is ", time.Now().Format(formatStr))
}
