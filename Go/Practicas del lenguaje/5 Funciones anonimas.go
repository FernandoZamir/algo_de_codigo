package main

import (
	"fmt"
	"time"
)

func main() {
	x := 10

	// Se deben usar cuadno su funcionalidad sea unica
	y := func() int {
		return x * 2
	}() // Con los parentesis llamamos a la función

	fmt.Println("El valor de la variable es:", y)

	// Uso de una gorutine
	c := make(chan int)
	go func() {
		fmt.Println("Iniciando función")
		time.Sleep(3 * time.Second)
		fmt.Println("Finalizando")
		c <- 1
	}()
	// Obtenemos el valor de la gorutine
	<-c
}
