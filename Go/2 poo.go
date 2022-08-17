package main

import "fmt"

// Creación de estructura
type Empleado struct {
	id        int
	nombre    string
	licenciaM bool
}

// Metodos asociados a la estructura

// Asignar id al empleado
func (e *Empleado) SetId(id int) {
	e.id = id
}

func (e *Empleado) SetName(name string) {
	e.nombre = name
}

// Devolver datos de la estructura
func (e *Empleado) GetId() int {
	return e.id
}

func (e *Empleado) GetName() string {
	return e.nombre
}

// Metodo 2, (Mejor opción, Crear)
func NuevoEmpleado(id int, nombre string, licenciam bool) *Empleado {
	return &Empleado{
		id:        id,
		nombre:    nombre,
		licenciaM: licenciam,
	}
}

func main() {
	// Por defecto asigna valores
	e := Empleado{} // Creamos la variable de tipo Empleado
	fmt.Printf("%v", e)

	// Asignamos los valores a la estructura
	e.id = 1
	e.nombre = "Soledad Cuadrado"
	fmt.Printf("%v", e)

	// Asignación por metodo
	e.SetId(5)
	e.SetName("Fernando Moreno")
	fmt.Printf("%v", e)

	// Returns
	fmt.Printf("%v", e.GetId())
	fmt.Printf("%v", e.GetName())

	e2 := Empleado{
		id:        1,
		nombre:    "Miriam Negrete Vertel",
		licenciaM: true,
	}

	fmt.Printf("%v\n", e2.GetName())

	// Nuevo empleado (Constructor)
	e3 := NuevoEmpleado(6, "Juan Moreno", true)
	fmt.Printf("%v\n", *e3, e3.GetName())
}
