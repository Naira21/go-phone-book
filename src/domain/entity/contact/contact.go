package contactModel

type Contact struct {
	userId  int
	name    string
	surname string
	phone   string
}

func Create(userId int, name, surname, phone string) *Contact {
	return &Contact{
		userId:  userId,
		name:    name,
		surname: surname,
		phone:   phone,
	}
}
