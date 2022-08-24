package main

import "testing" // Libreria estandar para test unitario

/*
	Se debe iniciar cad función de prueba con Test y el nombre de la función que se va a testear
	ej: TestCover

	Covertura para evalua test de codigo y saber que partes
	del codigo no tiene test

	Covertura
	go test -cover

	Windows --
	go test --coverprofile=coverage.out -> para uso de metricas

	Miramos la estructura funcional por código (Mejor forma)
	go tool cover --func=coverage.out

	Que código no hemos testeado atraves de HTML
	go tool cover --html=coverage.out

*/

func TestSum(t *testing.T) {
	// Test Básico
	total := Sum(94, 6)
	if total != 100 {
		t.Errorf("Fallo la prueba de suma, %d se esperaba %d", total, 100)
	}

	// Tabla de -> slice tipo struct para prueba mas completa
	// se puede realzar con un directorio de datos
	table := []struct {
		a int // P1
		b int // P2
		n int // Respuesta esperada
	}{
		// Inicializamos el struct
		{1, 2, 3},
		{2, 2, 4},
		{3, 5, 8},
	}

	// Valores _ wildcare
	for _, item := range table {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Fallo la prueba de suma, %d se esperaba %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	max := []struct {
		a int // P1
		b int // P2
		n int // Respuesta esperada
	}{
		{2, 1, 2},
		{2, 6, 6},
		{3, 5, 5},
	}

	for _, item := range max {
		maxR := Max(item.a, item.b)
		if maxR != item.n {
			t.Errorf("Fallo la prueba de maximo valor, %d se esperaba %d", maxR, item.n)
		}
	}
}
