package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine : i = ", i)
		time.Sleep(1 * time.Second)
	}
}


func main() {

	newTask()

	i := 0

	for {

		i++
		fmt.Println("main goroutine: i = ", i)
		time.Sleep(1 * time.Second)
	}
}