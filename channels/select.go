package channels

import (
	"fmt"
	"time"
)

func RunSelect() {
	// nil channel
	var nilChan chan struct{}
	// the following select will block for 1 second
	// because the nilChan is never read
	select {
	case <-nilChan: // infinite loop
	case <- time.After(1 * time.Second):
		fmt.Println("timed out")
		return
	default:
		fmt.Println("do something while waiting for timeout channel to be ready !!!")
	}
}