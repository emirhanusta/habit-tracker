package user

type CreateCommand struct {
	Username string
	Email    string
	Password string
}

type UpdateCommand struct {
	Id       string
	Username string
	Email    string
	Password string
}
