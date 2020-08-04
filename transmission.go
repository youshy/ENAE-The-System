package main

import (
	"fmt"
	"time"
)

var (
	transmission string = "Intercepting transmission\n"
	downloading  string = "Downloading:\n"
)

func loading() {
	printString(transmission, 52)
	printString(downloading, 52)
	totalTime := 2518 - 52 - 52 //
	progressBar(100, totalTime)
	cleanDisplay()
}

func progressBar(length, totalTime int) {
	timePerChar := totalTime / (length + 2) // we need to wrap [ ]
	fmt.Printf("[")
	for i := 0; i < length; i++ {
		fmt.Printf("|")
		time.Sleep(time.Millisecond * time.Duration(timePerChar))
	}
	fmt.Printf("]")
}

func solo() {

}
