package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	result := make(chan string, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			result <- strconv.Itoa(i)
		}(i)
	}
	// time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("begin channel")
	close(result)
	for i := range result {
		fmt.Println(i)
	}
}
