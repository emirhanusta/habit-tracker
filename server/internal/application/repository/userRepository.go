package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
	"habit-tracker/internal/domain"
)

type IUserRepository interface {
	GetUserByID(id string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	CreateUser(user domain.User) (*domain.User, error)
}

type userRepository struct {
	dbConn *pgx.Conn
}

func NewUserRepository(dbConn *pgx.Conn) IUserRepository {
	return &userRepository{
		dbConn: dbConn,
	}
}

func (u *userRepository) GetUserByID(id string) (*domain.User, error) {

	c := context.Background()
	getById := "SELECT * FROM users WHERE id = $1"

	queryRow := u.dbConn.QueryRow(c, getById, id)

	var user domain.User
	err := queryRow.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return &domain.User{}, errors.New(fmt.Sprint("Error getting user by id: ", err))
	}
	if err != nil {
		return &domain.User{}, errors.New(fmt.Sprint("Error getting user by id: ", err))
	}
	return &user, nil
}

func (u *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	c := context.Background()
	getByEmail := "SELECT * FROM USERS WHERE email = $1"

	queryRow := u.dbConn.QueryRow(c, getByEmail, email)
	var user domain.User
	err := queryRow.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return &domain.User{}, errors.New(fmt.Sprint("Error getting user by id: ", err))
	}
	if err != nil {
		return &domain.User{}, errors.New(fmt.Sprint("Error getting user by id: ", err))
	}
	return &user, nil

}

func (u *userRepository) CreateUser(user domain.User) (*domain.User, error) {
	c := context.Background()
	insertSql := "INSERT INTO users (id, email, password_hash, created_at) VALUES ($1, $2, $3, $4)"

	_, err := u.dbConn.Exec(c, insertSql, user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return &domain.User{}, errors.New(fmt.Sprint("Error creating user: ", err))
	}
	return &user, nil
}
