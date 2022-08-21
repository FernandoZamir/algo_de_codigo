package main

import "fmt"

/*
	Formatear resultados de funciones
*/

// Solución a la problematica de distintos valores en una función
// en otros lenguajes se puede utilizar objetos
func sum(value ...int) int {
	total := 0
	for _, num := range value {
		total += num
	}
	return total
}

// Solo impresión
func names(name ...string) {
	for _, n := range name {
		fmt.Println(n)
	}
}

func getValues(x int) (double int, triple int, cuadruple int) {

	double = 2 * x
	triple = 3 * x
	cuadruple = 4 * x
	return // Busca el valor de las variables a retornar y las devuelve
}

func suma(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Respuesta: ", suma(1, 4))
	fmt.Println("Respuesta: ", sum(1, 4, 8))
	names("Soledad", "Fernando", "Mario")
	fmt.Println(getValues(2))
}
