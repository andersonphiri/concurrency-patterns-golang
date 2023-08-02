/*
*
* lets suppose we have a queue of fixed length, say 2.
and we have more than 2 items to push onto the queue.
we want to push items as soon as there is space, that is, 
as soon as count is less than two
reasons could be, downstream consumer cannot keep up with upstream supplier
*/
package waitingforsignal

import (
	"fmt"
	"sync"
	"time"
)

const (
	MAX_ITEM_COUNT int = 2
)

type WaitForSignal struct { }

func (rcv WaitForSignal) RunWaitForSignal()  {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	consumer := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		// consume
		item := queue[0]
		fmt.Printf("value %v removed from queue!!!\n", item)
		queue = queue[1:]
		// then release lock
		c.L.Unlock()
		// then signal
		c.Signal()
	}
	// send to queue while consuming 
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == MAX_ITEM_COUNT {
			c.Wait()
		}
		fmt.Println("Adding to the queue")
		queue = append(queue, i)
		go consumer(1 * time.Second	)
		c.L.Unlock()
	}


}
