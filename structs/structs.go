package structs

type user struct {
	Name     string
	Age      uint8
	Address  address
	password string
}

type address struct {
	Street string
	Number uint8
	City   string
}

func CreateUser(name string, age uint8, address address, password string) user {
	user := user{
		Name:     name,
		Age:      age,
		Address:  address,
		password: password,
	}

	return user
}

func CreateAddress(street string, number uint8, city string) address {
	addr := address{
		Street: street,
		Number: number,
		City:   city,
	}

	return addr
}
