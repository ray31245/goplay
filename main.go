package main

import (
	"fmt"
	"time"
)

func main() {
	go warp()
	time.Sleep(time.Second * 5)
}

func warp() {
	go thefunc()
}

func thefunc() {
	fmt.Println("123")
	// panic("")
}
