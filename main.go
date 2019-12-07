package main

import (
	"fmt"

	// "github.com/ianhecker/golang-concurrency/concurrency"
	"github.com/ianhecker/golang-concurrency/addition"
	"github.com/ianhecker/golang-concurrency/read"
	"github.com/ianhecker/golang-concurrency/timer"
)

func main() {
	// runtime.GOMAXPROCS(1)

	routineList := read.ReadFile("./routines.data") //Read in lists from files
	integerList := read.ReadFile("./numbers.data")

	for _, routines := range *routineList { //For each routine count

		timer := timer.NewTimer()
		sum := addition.Add(*integerList, routines, timer)

		fmt.Printf("sum=%d at routines=%d took %dns=%dms\n", sum, routines, timer.Nanoseconds(), timer.Milliseconds())
	}
}
