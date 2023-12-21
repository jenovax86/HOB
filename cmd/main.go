package main

import (
	"fmt"
	"time"
)

func hobGame() {
	var count = 0
	for count < 50 {
		count++
		if count%5 != 0 {
			fmt.Println(count)
		} else {
			fmt.Println("HOB")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	hobGame()
}
