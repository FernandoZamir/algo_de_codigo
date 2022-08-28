package main

import (
	"fmt"
)

func main() {

	// Canal con bufferd
	c := make(chan int, 4)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
