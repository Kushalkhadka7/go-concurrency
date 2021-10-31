package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// It notifies the wait group that there is 1 go routine running.
	wg.Add(1)
	fmt.Println("Starting program execution...")
	go calculateSquares(55)

	fmt.Println("Ending program execution...")

	// It will halt the program exit until all the go routines are executed and, until all the go routines
	// signal done.
	wg.Wait()
}

func calculateSquares(number int) {
	sum := 0

	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number = number / 10
	}

	fmt.Println(sum)
	// It notifies the wait group that the execution of go routine is completed.
	defer wg.Done()

}
