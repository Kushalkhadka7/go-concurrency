# Go Concurrency

Go lang concurrency with channels and worker pools using sync.Pools.

## Concurrency

Concurrency is an ability of a program to do multiple things at the same time. This means a program that have two or more tasks that run individually of each other, at about the same time, but remain part of the same program.

In GO concurrency is the ability for functions to run independent of each other. A goroutine is a function that is capable of running concurrently with other functions.

```
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Stating program execution...")

    # Go routine.
	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("Hello from goroutine.")
	}()
	fmt.Println("Stopping program execution...")

    # This for statement is used to pause the program.
    # Since main func is also a go routine it will exit immediately without waiting for the go routine to complete.
	for {

	}
}

go run main.go

Output:

Stating program execution...
Stopping program execution...
Hello from goroutine.
```

## Channels

Channels are the medium for goroutines to communicate with one another. In channels data can be sent from one end and received from the other end using channels.

```
package main

import (
	"fmt"
)

func calculateSquares(number int, s chan int) {
	sum := 0

	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number = number / 10
	}

	s <- sum

    # Channels should be always closed and the closing should be always done by sending end instead of receiving end.
	defer close(s)
}

func main() {
	s := make(chan int)
	fmt.Println("Starting program execution...")
	go calculateSquares(55, s)

    # Channels are blocking so unless the data is received from channel `s`, the execution of the program won't move to next line.
	sum := <-s

	fmt.Println("Final output", sum)
	fmt.Println("Ending program execution...")
}

Output:

Starting program execution...
Final output 50
Ending program execution...

```

Here, chanel s is used to communicate between `calculateSquares` and `main` go routine.

## Wait groups.

To wait for multiple goroutines to finish, we can use a wait group. Wait groups comes under `sync` package.

```
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

```

## Sync pool

Frequent allocation and recycling of memory will cause a heavy burden to Garbage collector. `sync.Pool` caches objects that are not used temporarily and use them directly (without reallocation) when they are needed next time.

`Note`: example is in `worker_pools.go`
