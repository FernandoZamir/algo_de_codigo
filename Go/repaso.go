package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Variables
	//var x int = 7
	// y := 7
	// x

	// Conversiones
	valor, err := strconv.ParseInt("7", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(valor)
	}

	// Mapas -> Clave valor
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	// Slice -> Similar a un arreglo
	s := []int{2, 6, 4, 4}
	for index, value := range s {
		fmt.Println("Indice: ", index)
		fmt.Println("valor: ", value)
	}

	// Agregar elementos al Slice
	s = append(s, 45)

	for index, value := range s {
		fmt.Println("Indice: ", index)
		fmt.Println("valor: ", value)
	}

	// Canales para comunicar rutinas, tener encuenta siempre
	canal := make(chan int)
	// Creamos la rutina
	go holRutine(canal)
	<-canal

	// Punteros
	g := 20
	fmt.Println("valor G: ", g)
	h := &g // Refencia en memoria
	fmt.Println("valor H dirección: ", h)
	fmt.Println("valor H: ", *h)
}

// Simulando Go Rutines
func holRutine(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Simulación de rutina terminada.")
	c <- 1
}
