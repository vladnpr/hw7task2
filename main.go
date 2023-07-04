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
	var numsChan = make(chan int)
	var resultsChan = make(chan results)
	var nums []int

	for i := 0; i <= 100; i++ {
		go func() {
			select {
			case res := <-resultsChan:
				fmt.Printf("\nMax: %d, Min: %d", res.Max, res.Min)
			default:
				numsChan <- rand.Intn(100)
			}
		}()

		go func() {
			num := <-numsChan
			nums = append(nums, num)
			max, min := minMax(nums)
			res := results{
				Max: max,
				Min: min,
			}
			resultsChan <- res
		}()
	}

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
