package forselect

import (
	"fmt"
	"time"
)

// RunOrChannel demonstrate a situation where yiu have multiple channels
// and you would want to close as soon as any of
// these multiple channels are closed or are written to
// space complexity: floor(n/2)-goroutines
func RunOrChannel() {
	var or func(channels ...<- chan interface { }) <- chan interface {}
	or = func(channels ...<- chan interface { }) <- chan interface {} {
		switch len(channels) {
		case 0:
			return nil // since this is a recursive func, set terminating criteria
		case 1:// since this is a recursive func, set terminating criteria
		// if our variadic has only one item, then return this
			return channels[0] 
		}
		orDone := make(chan interface {})
		go func() { // main body of the function.  wait for messages on our channels without blocking
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <- channels[0]: 
				case <- channels[1]:
				}
			default:
				select {
					case <- channels[0]:
					case <- channels[1]:
					case <- channels[2]:
					case <- or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
	// test run 
	signal := func(duration time.Duration) <- chan interface {} {
		result := make(chan interface {})
		go func ()  {
			defer close(result)
			time.Sleep(duration)
		}()
		return result
	}
	start := time.Now()

	<- or(
		signal(2 * time.Hour),
		signal(5 * time.Minute),
		signal(3 * time.Second),
		signal(1 * time.Hour),
		signal(1 * time.Minute),
	)
	fmt.Printf("\ndone after %v\n", time.Since(start))
}