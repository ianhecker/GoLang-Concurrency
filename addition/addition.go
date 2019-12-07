package addition

import (
	"github.com/ianhecker/golang-concurrency/timer"
)

func Add(numbers []int, routines int, t *timer.Timer) int {
	t.Start()
	defer t.Stop()

	sums := make(chan int, len(numbers)/2) //create chan with buffer the size of sums of numbers
	subLen := len(numbers) / routines      //create length of subarrays

	for i := 0; i < routines; i++ { //create number = i of go routines for portions of array
		from := i * subLen //portion out array for routines
		until := (i + 1) * subLen

		if routines-i != 1 { //spawn routine with subportion
			go subAdd(numbers[from:until], sums)
		} else {
			go subAdd(numbers[from:len(numbers)], sums)
		}
	}

	var result int
	for i := 0; i < routines; i++ { //get each sum out of the sums channel
		select {
		case sum := <-sums:
			result += sum
		}
	}
	close(sums)
	return result
}

func subAdd(subarray []int, sums chan int) {
	var sum int
	for _, number := range subarray { //sum the subarray
		sum = sum + number
	}
	sums <- sum //inject sum into sums channel
}
