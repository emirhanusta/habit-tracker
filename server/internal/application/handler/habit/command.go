package habit

type CreateCommand struct {
	Name        string
	Description string
	UserId      string
}

type UpdateCommand struct {
	Id          string
	Name        string
	Description string
	UserId      string
}
