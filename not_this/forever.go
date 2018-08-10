package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		go func() {
			fmt.Println(time.Now())
		}()
	}
}
