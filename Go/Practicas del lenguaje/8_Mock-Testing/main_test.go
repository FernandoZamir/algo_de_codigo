package main

import (
	"testing"
)

func TestGetFullTimeEmployeeByID(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FulltimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id:       1,
						position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(id string) (Person, error) {
					return Person{
						Name: "Fernando Moreno",
						Age:  27,
						DNI:  "1",
					}, nil
				}
			},
			// Equivalencia de datos segun la estructura
			expectedEmployee: FulltimeEmployee{
				Person: Person{
					Name: "Fernando Moreno",
					Age:  27,
					DNI:  "1",
				},
				Employee: Employee{
					id:       1,
					position: "CEO",
				},
			},
		},
	}

	// Evitar para que se utilicen los mocks de manera permanente
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, itemTest := range table {
		itemTest.mockFunc()
		ft, err := GetFullTimeEmployeeByID(itemTest.id, itemTest.dni)

		if err != nil {
			t.Errorf("Error obteniendo datos del empleado")
		}

		// Segun el test verificar cada propiedad
		if ft.Age != itemTest.expectedEmployee.Age {
			t.Errorf("Error, se tiene %d y se espera %d", ft.Age, itemTest.expectedEmployee.Age)
		}
	}
	// Funciones originales a su posición
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI

}
