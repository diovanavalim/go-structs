package main

import (
	"fmt"
	"structures/structs"
)

func main() {
	var name string = "Sammy"
	var lastName string = "The Shark"
	var phoneNumber string = "48999454200"
	var age uint8 = 21
	var password string = "bolota"

	var street string = "Rua dos Perdizes"
	var number uint8 = 234
	var city string = "Morro Grande"

	var degree string = "Bsc"
	var college string = "UFSC"
	var course string = "Inf Sys"

	addr := structs.CreateAddress(street, number, city)

	usr := structs.CreateUser(name, age, addr, password)

	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
	// fmt.Println(usr.password) - O campo password não é visível fora do pacote structs

	p := structs.CreatePerson(name, lastName, phoneNumber, age)
	s := structs.CreateStudent(degree, college, course, p)

	// Imprimindo os campos da pessoa p no estudante s
	fmt.Println(s.Name)
	fmt.Println(s.PhoneNumber)
	fmt.Println(s.Age)
}
