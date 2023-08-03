package pools

import (
	"fmt"
	"sync"
)

type CreatePools struct{}

func (rcv CreatePools) RunCreatingPools() {
	var trackCount int64
	bufferPool := &sync.Pool{
		New: func() interface{} {
			// not necessary, but just to prove how memory efficient and effective 
			// are pools. in my test, for 1024*1024 workers, the trackCount was about 200 000
			// atomic.AddInt64(&trackCount, 1) 
			
			buffer := make([]byte, 1024)
			return &buffer
		},
	}
	// seed with 4 kb
	bufferPool.Put(bufferPool.New())
	bufferPool.Put(bufferPool.New())
	bufferPool.Put(bufferPool.New())
	bufferPool.Put(bufferPool.New())
	const numWorkers = 1024*1024 // 1024*
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	// then use the pool 
	for i := numWorkers; i > 0; i-- {
		go func ()  {
			defer wg.Done()
			buffer := bufferPool.Get().(*[]byte)
			// use this chunk for something more sensible
			// fmt.Printf(" %v\n", len(*buffer))
			defer bufferPool.Put(buffer) // return
		}()
	}
	wg.Wait()
	fmt.Printf(">> %d blocks were created, each of size 1024\n", trackCount)
	// trackCount should be less than 1024*1024 by a significant number
}