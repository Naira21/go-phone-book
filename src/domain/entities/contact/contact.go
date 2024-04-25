package contact_model

type Contact struct {
	userId   int
	name     string
	lastName string
	phone    string
}

func Create(userId int, name, lastName, phone string) *Contact {
	return &Contact{
		userId:   userId,
		name:     name,
		lastName: lastName,
		phone:    phone,
	}
}
