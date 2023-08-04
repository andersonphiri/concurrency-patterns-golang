package channels

import "fmt"

func RunExampleClearlyDefiningChannelOwnership() {
	chanOwner := func(bufferSize int) <-chan int {
		resultStream := make(chan int, bufferSize)
		go func(n int) {
			defer close(resultStream)
			for i := 0; i < n; i++ {
				resultStream <- i
			}
		}(bufferSize)
		return resultStream
	}
	intStream := chanOwner(5)
	for i := range intStream {
		fmt.Printf("Received: %v\n", i)
	}
	fmt.Println("done reading...")
}