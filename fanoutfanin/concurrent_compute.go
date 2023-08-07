package fanoutfanin

import "sync"

type ComputeFunc[TIn any, TOut any] func(done <-chan interface{},
	intputStream <-chan TIn) <-chan TOut



// ConcurrentComputes2 fans out workers and then merge their result into a single stream
func ConcurrentComputes[TIn any, TOut any](done <-chan interface{}, inputs <-chan TIn,
	parallelCount int, computer ComputeFunc[TIn, TOut]) <-chan TOut {
	workers := make(chan (<- chan TOut))
	var wg sync.WaitGroup 
	wg.Add(parallelCount)
	doWork := func (aggregate chan (<-chan TOut))  {
		select {
		case  <-done:
			return
		case aggregate <- computer(done, inputs):
		}
	}
	for i := 0; i < parallelCount; i++ {
		go doWork(workers)
	}
	go func(){
		wg.Wait()
		close(workers)
	}()
	return fanInStream[TOut](done, workers)

}

// FanIn takes in an array of readonly channels to produce a single readoly channel
func FanIn[T any](done <-chan interface{}, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	muxedStream := make(chan T)
	muxer := func (channel <-chan T) {
		defer wg.Done()
		for item := range channel {
			select {
			case <- done:
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
	go func () { 
		wg.Wait()
		close(muxedStream)
	 }()
	return muxedStream
}

func fanInStream[T any](done <-chan interface{}, channels chan (<- chan T)) <-chan T {
	var wg sync.WaitGroup
	muxedStream := make(chan T)
	muxer := func (channel <-chan T) {
		defer wg.Done()
		for item := range channel {
			select {
			case <- done:
				return
			case muxedStream <- item:
			}
		}
	}
	wg.Add(len(channels))
	for channel := range channels {
		go muxer(channel)
	}
	// wait for all reads to complete
	go func () { 
		wg.Wait()
		close(muxedStream)
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