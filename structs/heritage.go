package structs

type person struct {
	Name        string
	LastName    string
	PhoneNumber string
	Age         uint8
}

type student struct {
	person
	Degree  string
	College string
	Course  string
}

func CreatePerson(name, lastName, phoneNumber string, age uint8) person {
	p := person{
		name,
		lastName,
		phoneNumber,
		age,
	}

	return p
}

func CreateStudent(degree, college, course string, person person) student {
	s := student{
		person,
		degree,
		college,
		course,
	}

	return s
}
