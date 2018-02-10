package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	numbers := []uint32{125424, 235, 1241877, 32135647, 6544774}
	wg.Add(len(numbers))

	fmt.Println("Comienza el proceso...")
	for _, value := range numbers {
		go func(a uint32) {
			defer wg.Done()
			fmt.Println(a, EsPrimo(a))
		}(value)
	}
	wg.Wait()
	fmt.Println("Termino el proceso")
}

func EsPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}
