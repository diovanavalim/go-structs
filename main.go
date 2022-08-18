package main

import (
	"fmt"
	"structures/structs"
)

func main() {
	var name string = "Sammy The Shark"
	var age uint8 = 21
	var password string = "bolota"

	var street string = "Rua dos Perdizes"
	var number uint8 = 234
	var city string = "Morro Grande"

	addr := structs.CreateAddress(street, number, city)

	usr := structs.CreateUser(name, age, addr, password)

	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
	// fmt.Println(usr.password) - O campo password não é visível fora do pacote structs
}
