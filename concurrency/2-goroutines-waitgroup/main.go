package main

import (
	"sync"
	"time"
)

var wg *sync.WaitGroup

func init() {
	wg = new(sync.WaitGroup)
}

// WaitGroup --> State/Counter
// wg.Add(1)
// wg.Done() --> decrease the counter
// wg.Wait() --> This waits untilthe state becomes zero

func main() {
	defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("Hello World")
		defer time.Sleep(time.Millisecond * 10)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sum := func(nums ...int) int {
			sum := 0
			for _, n := range nums {
				sum += n
			}
			return sum
		}(12, 123, 14, 123, 213, 213, 12312, 3)
		println(sum)
	}()
	println("Hello end of main")
}

// Main is also a goroutine
// No gorutine waits for other goroution to complete its execution
// There is no concept of parent and child go routines
// cannot guarantee the order of execution
// just use go keyword to run a block of code as a goroutine
