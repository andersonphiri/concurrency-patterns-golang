package fanoutfanin

import (
	"fmt"
	"sync"
	"time"
)

type ComputeFunc[TIn any, TOut any] func(done <-chan interface{},
	intputStream <-chan TIn) <-chan TOut

// ConcurrentComputes fans out workers and then merge their result into a single stream
func ConcurrentComputes[TIn any, TOut any](done <-chan interface{}, inputs <-chan TIn,
	parallelCount int, computer ComputeFunc[TIn, TOut]) <-chan TOut {
	workers := make(chan (<-chan TOut))
	var wg sync.WaitGroup
	wg.Add(parallelCount)
	doWork := func(aggregate chan (<-chan TOut)) {
		select {
		case <-done:
			return
		case aggregate <- computer(done, inputs):
		}
	}
	for i := 0; i < parallelCount; i++ {
		go doWork(workers)
	}
	go func() {
		wg.Wait()
		close(workers)
	}()
	return fanInStream[TOut](done, workers)

}

// FanIn takes in an array of readonly channels to produce a single readoly channel
func FanIn[T any](done <-chan interface{}, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	muxedStream := make(chan T)
	muxer := func(channel <-chan T) {
		defer wg.Done()
		for item := range channel {
			select {
			case <-done:
				return
			case muxedStream <- item:
			}
		}
	}
	wg.Add(len(channels))
	for _, channel := range channels {
		go muxer(channel)
	}
	// wait for all reads to complete
	go func() {
		wg.Wait()
		close(muxedStream)
	}()
	return muxedStream
}

func fanInStream[T any](done <-chan interface{}, channels chan (<-chan T)) <-chan T {
	//var wg sync.WaitGroup
	muxedStream := make(chan T)
	
	// wait for all reads to complete
	go func() {
		defer close(muxedStream)
		muxer := func(channel <-chan T) {
			for item := range channel {
				select {
				case <-done:
					return
				case muxedStream <- item:
				}
			}
		}
		// wg.Add(len(channels))
		for channel := range channels {
			go muxer(channel)
		}
	}()
	return muxedStream
}

// ConcurrentComputesUsingSliceWorkers fans out workers and then merge their result into a single stream
func ConcurrentComputesUsingSliceWorkers[TIn any, TOut any](done <-chan interface{}, inputs <-chan TIn,
	parallelCount int, computer ComputeFunc[TIn, TOut]) <-chan TOut {
	workers := make([]<-chan TOut, 0, parallelCount)
	for i := 0; i < parallelCount; i++ {

		workers = append(workers, computer(done, inputs)) // this may require optimisation if compute is a long running operation
	}
	return FanIn[TOut](done, workers...)

}

type IsPrimeResult struct {
	Num     int
	IsPrime bool
}

// RunConcurrentComputes brings all ideas into a sing test run
func RunConcurrentComputesPrimeNumbers(start, end, parallesimFactor uint64) {
	checkIfPrime := func(num int) bool {
		n := num
		if num < 0 {
			n = -1 * n
		}
		if n == 1 {
			return false
		}
		if n == 0 || n == 3 {
			return true
		}
		if n%2 == 0 || n%3 == 0 {
			return false
		}
		for i := 5; i*i <= n; i = i + 6 {
			if n%i == 0 || n%(i+2) == 0 {
				return false
			}
		}
		return true
	}
	var computePrime ComputeFunc[int, IsPrimeResult] = func(done <-chan interface{},
		intputStream <-chan int) <-chan IsPrimeResult {
		result := make(chan IsPrimeResult)
		go func() {
			defer close(result)
			for num := range intputStream {
				select {
				case <-done:
					return
				case result <- IsPrimeResult{Num: num, IsPrime: checkIfPrime(num)}:
				}
			}

		}()
		return result

	}
	generateRandomStream := func(done <-chan interface{}, s,n int) <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for i := s; i <= n; i++ {
				select {
				case <-done:
					return
				case stream <- i:
				}
			}

		}()
		return stream
	}
	done := make(chan interface{})
	defer close(done)
	inputStream := generateRandomStream(done, int(start), int(end))
	startTime := time.Now()
	fmt.Println("beginning computations... Please wait")
	computations := ConcurrentComputesUsingSliceWorkers [int, IsPrimeResult](done, //ConcurrentComputes
		inputStream,
		int(parallesimFactor), computePrime)
	fmt.Println("... Please wait")
	count := 0
	for result := range computations {
		// fmt.Printf("result: %v\n", result)
		if result.IsPrime {
			count++
		}
	}

	fmt.Printf("There are %d prime numbers between %d and %d inclusive\n", count,start,end)
	fmt.Printf("The computation took: %v with parallelism factor of channels set to: %v\n",time.Since(startTime),
	 parallesimFactor)
	fmt.Println("completed...")
}
