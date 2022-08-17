package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

type Empleado struct {
	idEmpleado int
}

// Composición de herencia
type EmpleadoTiempoCompleto struct {
	Persona
	Empleado
	fechaFin string
}

func (etc EmpleadoTiempoCompleto) Mensaje() string {
	return "Empleado tiempo completo" //fmt.Println("Empleado timepo completo \n: ", etc.nombre)
}

type EmpleadoTiempoTemporal struct {
	Persona
	Empleado
	tiempo int
}

func (etc EmpleadoTiempoTemporal) Mensaje() string {
	return "Empleado tiempo temporal" //fmt.Println("Empleado timepo completo \n: ", etc.nombre)
}

type MensajeI interface {
	Mensaje() string
}

func Mensaje(m MensajeI) {
	fmt.Println(m.Mensaje())
}

// func Mensaje(p Persona) {
// 	fmt.Printf("%s con una edad de %d \n: ", p.nombre, p.edad)
// }

func main() {
	ftEmpleado := EmpleadoTiempoCompleto{}
	ftEmpleado.nombre = "Fernando Moreno OViedo"
	ftEmpleado.edad = 27
	ftEmpleado.idEmpleado = 1

	fmt.Println("Datos del empleado: ", ftEmpleado)

	// La herencia directamenta no funciona, la solución son las interfaces
	//Mensaje(ftEmpleado)

	eTiempoTemporal := EmpleadoTiempoTemporal{}

	// Uso de interfaces como patron polimorfico
	Mensaje(eTiempoTemporal)
	Mensaje(ftEmpleado)
}
