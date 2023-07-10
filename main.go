package main

import (
	"fmt"
	"math/rand"
	"time"
)

type results struct {
	Max int
	Min int
}

func main() {
	var numsChan = make(chan []int)
	var resultsChan = make(chan results)

	go func() {

		var randNums []int

		for i := 0; i < 10; i++ {
			randNums = append(randNums, rand.Intn(100))
		}

		numsChan <- randNums
		time.Sleep(10 * time.Millisecond)

		res := <-resultsChan
		fmt.Printf("\nMax: %d, Min: %d", res.Max, res.Min)
	}()

	go func() {
		nums := <-numsChan
		max, min := minMax(nums)
		res := results{
			Max: max,
			Min: min,
		}
		resultsChan <- res
	}()

	time.Sleep(2 * time.Second)
}

func minMax(nums []int) (max int, min int) {
	max = nums[0]
	min = nums[0]
	for _, value := range nums {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	return max, min
}
