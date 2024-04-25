package user_model

type User struct {
	id       int
	name     string
	lastName string
}

func Create(name, lastName string) *User {
	return &User{
		name:     name,
		lastName: lastName,
	}
}
