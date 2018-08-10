package main

import "fmt"

func main() {
	maxLimit := 10
	longList := [1000]int{}
	sem := make(chan bool, maxLimit)

	for _, i := range longList {
		sem <- true
		go func(work int) {
			defer func() { <-sem }()
			doStuff(work)
		}(i)
	}

	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}

func doStuff(work int) {
	fmt.Printf("stuff: %d\n", work)
}
