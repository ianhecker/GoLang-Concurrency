package main

import (
	"github.com/ianhecker/golang-concurrency/addition"
	"github.com/ianhecker/golang-concurrency/plot"
	// "github.com/ianhecker/golang-concurrency/random"
	"github.com/ianhecker/golang-concurrency/read"
	"github.com/ianhecker/golang-concurrency/timer"
)

func main() {

	routineList := read.ReadFile("./routines.data") //Read in lists from files
	integerList := read.ReadFile("./million_numbers.data")
	scatterPlot := plot.MakeScatterPlot(len(*routineList))

	for index, routines := range *routineList { //For each routine count

		timer := timer.NewTimer() //Timer for duration of go routines
		addition.Add(*integerList, routines, timer)

		scatterPlot.AddPoint(routines, timer.Nanoseconds(), index) //plot go routine count vs time
	}

	scatterPlot.Plot("Go_Routines_vs_Time_Nanoseconds") //plot all points in a png
}
