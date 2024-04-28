package userEntity

type User struct {
	id      int
	name    string
	surname string
}

func Create(name, surname string) *User {
	return &User{
		name:    name,
		surname: surname,
	}
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Surname() string {
	return u.surname
}

func (u *User) ID() int {
	return u.id
}
