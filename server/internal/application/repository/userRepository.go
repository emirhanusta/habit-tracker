package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
	"habit-tracker/internal/domain"
)

type IUserRepository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	SaveUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
	dbConn *pgx.Conn
}

func NewUserRepository(dbConn *pgx.Conn) IUserRepository {
	return &userRepository{
		dbConn: dbConn,
	}
}

func (u *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {

	if user.Email != "" && user.Username != "" && user.Password != "" {
		updateSql := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
		_, err := u.dbConn.Exec(ctx, updateSql, user.Username, user.Email, user.Password, user.Id)
		if err != nil {
			return errors.New(fmt.Sprint("Error updating user: ", err))
		}
		return nil
	} else {
		return errors.New(fmt.Sprint("Error updating user: ", "Missing fields"))
	}
}

func (u *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {

	getAll := `Select * FROM users`

	rows, err := u.dbConn.Query(ctx, getAll)
	if err != nil {
		fmt.Printf("userRepository.GetAll error: %s\n", err)
		return nil, err
	}

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			fmt.Printf("userRepository.GetAll error: %s\n", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
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

func (u *userRepository) DeleteUser(ctx context.Context, id string) error {
	deleteSql := `DELETE FROM users WHERE id = $1`

	_, err := u.dbConn.Exec(ctx, deleteSql, id)
	if err != nil {
		return errors.New(fmt.Sprint("Error deleting user: ", err))
	}
	return nil
}
