package main

import (
	"fmt"
	"sync"
	"time"
)

// Waitgroup

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Iniciado %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado")
}
