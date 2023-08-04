package pipelines

import "fmt"

func RunSimplePipelineSpaceInEfficient() {
	multiply := func(nums []int, multiplier int) []int {
		result := make([]int, 0, len(nums))
		for _, num := range nums {
			result = append(result, num*multiplier)
		}
		return result
	}
	add := func(nums []int, additive int) []int {
		result := make([]int, 0, len(nums))
		for _, num := range nums {
			result = append(result, num+additive)
		}
		return result
	}
	// the next lines are inefficient, they create two arrays of same size as input
	inputs := []int{1, 2, 3, 4}
	for _, finalItem := range multiply(add(inputs, 10), 20) {
		fmt.Printf("%v\n", finalItem)
	}
}
func RunSimplePipelineSpaceEfficient() {
	multiplyOne := func(num int, multiplier int) int {
		return num * multiplier
	}
	addOne := func(num int, additive int) int {
		return num+additive
	}
	// the next lines are inefficient, they create two arrays of same size as input
	inputs := []int{1, 2, 3, 4}
	for _, item := range inputs {
		fmt.Printf("%v\n", multiplyOne(addOne(item , 10), 20))
	}
}
func RunSimplePipelineUsingChannels() {
	generator := func(done <- chan interface {}, inputs ...int) <- chan int {
		istream := make(chan int)
		go func() {
			defer close(istream)
			for _, input := range inputs {
				select {
				case <- done:
					return 
				case istream <- input:
				}
			}
		}()
		return istream
	}
	multiply := func(done <- chan interface {}, inputs <- chan int, multiplier int) <- chan int {
		multipliedStream := make(chan int)
		go func ()  {
			defer close(multipliedStream)
			for input := range inputs {
				select {
				case <- done:
				case multipliedStream <- input * multiplier:

				}
			}
		}()
		return multipliedStream
	}
	add := func(done <- chan interface {}, inputs <- chan int, additive int) <- chan int {
		addedStream := make(chan int)
		go func ()  {
			defer close(addedStream)
			for input := range inputs {
				select {
				case <- done:
				case addedStream <- input + additive:

				}
			}
		}()
		return addedStream
	}
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1,2,3,4)
	pipeline := multiply(done, add(done, intStream, 10), 20)
	for value := range pipeline {
		fmt.Println(value)
	}

}