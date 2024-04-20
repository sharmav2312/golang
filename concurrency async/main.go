package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printNumbers() {
	defer wg.Done()
	n := []int{1, 2, 3, 4, 5}
	for i := range n {
		time.Sleep(time.Second)
		fmt.Println(n[i])
	}
}

func main() {
	now := time.Now()
	fmt.Println("Start")
	wg.Add(2)
	go printNumbers()
	go printNumbers()

	wg.Wait()
	fmt.Println("End")
	fmt.Println("elapsed: ", time.Since(now))
}
