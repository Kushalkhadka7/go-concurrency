package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomNo int
}

type CalculationResult struct {
	job         Job
	sumOfDigits int
}

func main() {
	startTime := time.Now()

	noOfJobs := 100
	noOfWorkers := 10

	jobs := make(chan Job)
	done := make(chan bool)
	result := make(chan CalculationResult)

	go createJobs(noOfJobs, jobs)

	go createWorker(noOfWorkers, jobs, result)

	go calculateResult(result, done)

	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff*time.Second, "seconds")
}

// Receives the result from the CalculationResult channel and prints it.
// Once all the result is printed from all workers, it will notify the program that the calculation is completed via done channel.
func calculateResult(res chan CalculationResult, done chan bool) {
	for result := range res {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomNo, result.sumOfDigits)
	}

	done <- true
}

// Create workers to do the job, for each worker 1 go routine is created .
// Workers complete the job and send the result to CalculationResult channel.
func createWorker(workers int, jobs <-chan Job, res chan CalculationResult) {
	var wg sync.WaitGroup

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			for job := range jobs {
				r := CalculationResult{job, calculate(job.randomNo)}

				res <- r
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	close(res)
}

// Create jobs and send the job to job channel.
func createJobs(noOfJobs int, jobs chan Job) {
	for i := 0; i < noOfJobs; i++ {
		randoNo := rand.Intn(999)
		job := Job{i, randoNo}

		jobs <- job
	}

	close(jobs)
}

// Calculate the sum of the given number and return it.
func calculate(number int) int {
	sum := 0
	no := number

	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	time.Sleep(2 * time.Second)

	return sum
}
