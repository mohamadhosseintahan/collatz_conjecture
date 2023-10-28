package main

import (
	"fmt"
	"sync"
)

func Collatz(n, counter uint) (uint, uint) {
	if n == 1 {
		return 1, counter
	}
	counter++
	if n%2 == 1 {
		return Collatz((n*3)+1, counter)
	} else {
		return Collatz(n/2, counter)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go func (n uint) {
			defer wg.Done()
			var counter uint = 0
			_, b := Collatz(n, counter)
			fmt.Printf("%d ----- %d\n", n, b)
		}(uint(i))
	}
	wg.Wait()
}
