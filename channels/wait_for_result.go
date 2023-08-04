package channels

import (
	"log"
	"time"
)

// WaitForResultFromAnother demonstrates an example of
// one go routine do some other work while waiting for another
// go routine to report result
// this is achieved using a select-for / for-select loop
func WaitForResultFromAnother(doneFunc func ()) {
	done := make(chan interface{})
	go func () {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0
	for {
		select {
		case <- done:
			if doneFunc != nil {
				doneFunc()
			}
			log.Printf("the variable has been incremented to : %v\n",workCounter)
			return
		default:
			// no op
		}
		// simulate do work 
		workCounter++
		time.Sleep(1 * time.Second)
	}

}