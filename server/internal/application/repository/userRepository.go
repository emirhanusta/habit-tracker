package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
	"habit-tracker/internal/domain"
)

type IUserRepository interface {
	GetById(ctx context.Context, id string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	SaveUser(ctx context.Context, user *domain.User) error
}

type userRepository struct {
	dbConn *pgx.Conn
}

func NewUserRepository(dbConn *pgx.Conn) IUserRepository {
	return &userRepository{
		dbConn: dbConn,
	}
}

func (u *userRepository) GetById(ctx context.Context, id string) (*domain.User, error) {

	getById := `SELECT * FROM users WHERE id = $1`

	queryRow := u.dbConn.QueryRow(ctx, getById, id)

	var user domain.User
	err := queryRow.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		fmt.Printf("Error getting user by id: %s\n", err)
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	getByEmail := `SELECT * FROM USERS WHERE email = $1`

	queryRow := u.dbConn.QueryRow(ctx, getByEmail, email)
	var user domain.User
	err := queryRow.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return &domain.User{}, errors.New(fmt.Sprint("Error getting user by id ", err))
	}
	if err != nil {
		return &domain.User{}, errors.New(fmt.Sprint("Error getting user by id ", err))
	}
	return &user, nil

}

func (u *userRepository) SaveUser(ctx context.Context, user *domain.User) error {

	insertSql := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

	_, err := u.dbConn.Exec(ctx, insertSql, user.Username, &user.Email, &user.Password)
	if err != nil {
		return errors.New(fmt.Sprint("Error creating user: ", err))
	}
	return nil
}
