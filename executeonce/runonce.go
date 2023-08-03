package executeonce

import (
	"fmt"
	"sync"
)

type ExecuteOnce struct {
}

type DoItem struct {
	sync.Once
}
// func (rcv ExecuteOnce) RunExecuteOnce_Two() {
// 	fmt.Println("running sync.Once example two(this run with data races)...")
// 	var count int
// 	increment := func() {
// 		count++
// 	}
// 	items := []DoItem{
// 		{sync.Once{}},{sync.Once{}},{sync.Once{}},{sync.Once{}},
// 		{sync.Once{}},{sync.Once{}},{sync.Once{}},{sync.Once{}},
// 		{sync.Once{}},{sync.Once{}},{sync.Once{}},{sync.Once{}},
// 		{sync.Once{}},{sync.Once{}},{sync.Once{}},{sync.Once{}},
// 		{sync.Once{}},{sync.Once{}},{sync.Once{}},{sync.Once{}},
// 	}
// 	var increments sync.WaitGroup 
// 	increments.Add(len(items))
// 	for _, item := range items {
// 		go func(i DoItem){
// 			defer increments.Done()
// 		  i.Do(increment)
// 		}(item)
// 	}
// 	increments.Wait()
// 	fmt.Printf("this time, the value of count is: %v\n", count)

// }

func (rcv ExecuteOnce) RunExecuteOnce() {
	fmt.Println("running sync.Once example...")
	var count int
	increment := func() {
		count++
	}
	var once sync.Once
	var increments sync.WaitGroup 
	increments.Add(20)
	for i := 0; i < 20; i++ {
		go func ()  {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("instead of 20, the value of count is: %v\n", count)
}