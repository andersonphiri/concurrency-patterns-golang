/*
*
imagine we have a button.
we want to execute a number of functions when that button is clicked
lets call these, handlers
*/
package waitingforsignal

import (
	"fmt"
	"sync"
)

type Button struct {
	clicked *sync.Cond
}

type BroadcastExample struct {

}

func (rcv BroadcastExample) RunBroadcast() {
	fmt.Println("running cond broadcast example...\n")
	var button = Button{clicked: sync.NewCond(&sync.Mutex{}) }
	subscribe := func (cond *sync.Cond, handle func())  {
		var running sync.WaitGroup
		running.Add(1)
		go func ()  {
			running.Done()
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			handle()
		}()
		running.Wait()
	}
	var clickRegistered sync.WaitGroup 
	clickRegistered.Add(3)
	subscribe(button.clicked, func() {
		fmt.Println("I am subscriber number one")
		clickRegistered.Done()
	})
	subscribe(button.clicked, func() {
		fmt.Println("I am subscriber number two(2)")
		clickRegistered.Done()
	})
	subscribe(button.clicked, func() {
		fmt.Println("I am subscriber number three(3)")
		clickRegistered.Done()
	})
	// then broadcast
	button.clicked.Broadcast()
	clickRegistered.Wait()
	fmt.Println("all registered subscribers successfully executed")
}


