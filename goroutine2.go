package main

import (
	"runtime"
	"time"
	"fmt"
)

func main() {

	go func() {
		defer fmt.Println("A.defer")
	
		func() {
			defer fmt.Println("B.defer")

			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}