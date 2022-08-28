package main

/*
	Multiplexación
*/
import (
	"fmt"
	"time"
)

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	for i := 0; i < 2; i++ {
		select { // Para utilizar con diferentes canales
		case chan1 := <-c1:
			fmt.Println(chan1)

		case chan2 := <-c2:
			fmt.Println(chan2)
		}
	}
}
