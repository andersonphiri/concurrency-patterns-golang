package forselect

import (
	"fmt"
	"math/rand"
	"time"
)

// RunDoneChannel eliminates a go-routine leak by using a 'Done' channel
// the done channel will be cloased by parent go routing or the forking goro
func RunDoneChannel() {
	doWork := func(done <-chan interface{}, strings <-chan string) <- chan interface{} {
		terminated := make(chan interface{})
		go func ()  {
			defer fmt.Println("doing work completed, exited")
			defer close(terminated)
			for {
				select {
				case s := <- strings: // here select will block indefinitely since the channel is never ready
				// at least in this setup
					fmt.Println(s)
				case <- done:
					return 
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{ })
	terminated := doWork(done, nil )
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("cancelling doWork goroutine...")
		close(done) // signal to child routine to exit
	}()
	<- terminated 
	fmt.Println("Done the saga")
	
}

func RunTellProducerToStop() {
	newRandStream := func(done <- chan interface{}) <- chan int {
		result := make(chan int)
		go func ()  {
			defer fmt.Printf("closing producer as signaled by downstream consumer which happens to be the parent goroutine\n")
			defer close(result)
			for {
				select {
				case result <- rand.Int():
				case <- done:
					return // will call all defer blocks
				}
			}
		}()
		return result
	}
	done := make(chan interface{})
	intStream := newRandStream(done)
	// read first three nums
	for i := 0; i < 3; i++ {
		fmt.Printf("stream count(%d) -> %d: %d\n",len(intStream), i, <-intStream)
	}
	// then signal producer to stop
	close(done)
	fmt.Printf("\nthe stream had %d items after producer was signaled to stop\n", len(intStream))

}