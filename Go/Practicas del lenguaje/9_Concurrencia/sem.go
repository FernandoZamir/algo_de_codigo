package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		c <- 1    // el canal envia 1, en un cupo
		wg.Add(1) // Añado al contador
		go doSomething(i, &wg, c)
	}

	// Bloqueo
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done() // Finalizo la rutina
	fmt.Printf("Id %d iniciado\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finalizado\n", i)
	<-c // Libera el espacio, esperando su respuesta
}
