package confinement

import (
	"bytes"
	"fmt"
	"sync"
)

// RunBasicLexicalConfinement achieves confinement by specifying whether the channel you are
// consuming from is read only or write only
func RunBasicLexicalConfinement() {
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
	
	consumerFunc := func (ch <- chan int)  {
		for i := range ch {
			fmt.Printf("Received: %v\n", i)
		}
		fmt.Println("done reading using lexical confinement...")
	}
	intStream := chanOwner(5)
	consumerFunc(intStream)
	
}

func OperateOnNonThreadSafeTypeConcurrently() {
	printData := func (wg *sync.WaitGroup, data []byte)  {
		defer wg.Done()
		var buffer bytes.Buffer
		for _, char := range data {
			fmt.Fprintf(&buffer,"%c", char)
		}
		fmt.Println(buffer.String())
	}
	var wg sync.WaitGroup 
	wg.Add(2)
	data := []byte("Anderson Phiri")
	// go printData(&wg, data[:8])
	// go printData(&wg, data[9:])
	go printData(&wg, data)
	go printData(&wg, data)
	wg.Wait()
}