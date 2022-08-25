package main

import "time"

// go mod init github.com/FernandoZamir/8_Mock-Testing

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	id       int
	position string
}

type FulltimeEmployee struct {
	Employee
	Person
}

/*
	Modificamos las funciones para ser testeadas
	var nom_fun = func nom_fun
*/

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(3 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(3 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeByID(id int, dni string) (FulltimeEmployee, error) {
	var ftEmployee FulltimeEmployee

	Emp, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = Emp

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}
