package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{} //A WaitGroup waits for a collection of goroutines to finish

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("One")       // add one in score
		score = append(score, 1) // append in score
		wg.Done()                // Done decrements the WaitGroup counter by one
	}(wg) // &sync.WaitGroup{}

	//wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two") // add two in score
		score = append(score, 2)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("tree") // add tree in score
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait() // Wait blocks until the WaitGroup counter is zero
	fmt.Println(score)
	//output will be always random order
}
