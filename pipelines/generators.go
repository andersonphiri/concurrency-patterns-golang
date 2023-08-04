package pipelines

import "fmt"

func Generate[T any](done <-chan interface{}, inputs ...T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for _, value := range inputs {
			select {
			case <-done:
				return
			case stream <- value:
			}
		}
	}()
	return stream
}

func GenerateFunc[TOut any, TIn any](done <-chan interface{}, 
	supplierFunc func(TIn)TOut, inputs ...TIn) <-chan TOut {
	stream := make(chan TOut)
	go func() {
		defer close(stream)
		for _, value := range inputs {
			select {
			case <-done:
				return
			case stream <- supplierFunc(value):
			}
		}
	}()
	return stream
}

func GenerateFuncIf[TOut any, TIn any](done <-chan interface{}, 
	supplierFunc func(TIn)TOut, predicate func(TIn,TOut)bool, inputs ...TIn) <-chan TOut {
	stream := make(chan TOut)
	go func() {
		defer close(stream)
		for _, value := range inputs {
			newItem := supplierFunc(value)
			if !predicate(value, newItem){
				continue
			}
			select {
			case <-done:
				return
			case stream <- newItem:
			}
		}
	}()
	return stream
}
func GenerateFuncIfIn[TOut any, TIn any](done <-chan interface{}, 
	supplierFunc func(TIn)TOut, predicate func(TIn)bool, inputs ...TIn) <-chan TOut {
	stream := make(chan TOut)
	go func() {
		defer close(stream)
		for _, value := range inputs {
			if !predicate(value){
				continue
			}
			select {
			case <-done:
				return
			case stream <- supplierFunc(value):
			}
		}
	}()
	return stream
}
func GenerateFuncIfOut[TOut any, TIn any](done <-chan interface{}, 
	supplierFunc func(TIn)TOut, predicate func(TOut)bool, inputs ...TIn) <-chan TOut {
	stream := make(chan TOut)
	go func() {
		defer close(stream)
		for _, value := range inputs {
			newItem := supplierFunc(value)
			if !predicate(newItem){
				continue
			}
			select {
			case <-done:
				return
			case stream <- newItem:
			}
		}
	}()
	return stream
}

func Repeat[T any](done <-chan interface{}, unitValue   T, repeatCount int) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for i:= 0; i <repeatCount; i++  {
			select {
			case <-done:
				return
			case stream <- unitValue:
			}
		}
	}()
	return stream
}

func TakeN[T any](done <-chan interface{}, input <-chan T, n int) <-chan T {
	takeStream := make(chan T)
	go func() {
		defer close(takeStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-input:
			}
		}
	}()
	return takeStream
}

func RunUseGenerateAndTake() {
	done := make(chan interface{})
	defer close(done)
	pipeline := TakeN[int](done, Repeat[int](done, 2, 10), 10)
	// must print 1s ten times
	for num := range pipeline {
		fmt.Printf("%v ", num)
	}
	fmt.Println()
}