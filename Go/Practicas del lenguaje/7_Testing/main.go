package main

func Sum(x, y int) int {
	return x + y
}

func Max(x, y int) int {

	if x > y {
		return x
	}

	return y
}

// Fibo recursivo
func Frecursivo(n int) int {
	if n <= 1 {
		return n
	}

	return Frecursivo(n-1) + Frecursivo(n-2)
}
