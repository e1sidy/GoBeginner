package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	// Formating time
	fmt.Printf("%d.%d.%d\n", t.Day(), t.Month(), t.Year())

	// Adding a week to the current time
	week := time.Duration(7 * 24 * 60 * 60 * 1e9) // expressing week in nanoseconds
	nextWeek := t.Add(week)
	fmt.Println(nextWeek)

	// Fomating using time.Format()
	// TODO : see into significance of  02 Jan 2006 15:04
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	fmt.Println(t.Format("2006 01 02"))

}
