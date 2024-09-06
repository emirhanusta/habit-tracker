package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"habit-tracker/internal/domain"
)

type IHabitRepository interface {
	GetAllByUserId(ctx context.Context, id string) ([]domain.Habit, error)
	GetById(ctx context.Context, id string) (*domain.Habit, error)
	Save(ctx context.Context, habit *domain.Habit) error
	Update(ctx context.Context, habit *domain.Habit) error
	Delete(ctx context.Context, id string) error
}

type habitRepository struct {
	dbConn         *pgx.Conn
	userRepository IUserRepository
}

func NewHabitRepository(dbConn *pgx.Conn) IHabitRepository {
	return &habitRepository{
		dbConn: dbConn,
	}
}

func (h *habitRepository) GetAllByUserId(ctx context.Context, id string) ([]domain.Habit, error) {

	getAllByUserId := `Select * FROM habits WHERE user_id = $1`

	rows, err := h.dbConn.Query(ctx, getAllByUserId, id)
	if err != nil {
		fmt.Printf("habitRepository.GetAllByUserId error: %s\n", err)
		return nil, err
	}

	var habits []domain.Habit

	for rows.Next() {
		var habit domain.Habit
		err = rows.Scan(&habit.Id, &habit.Name, &habit.Description, &habit.UserId)
		if err != nil {
			fmt.Printf("habitRepository.GetAllByUserId error: %s\n", err)
			return nil, err
		}
		habits = append(habits, habit)
	}

	return habits, nil
}

func (h *habitRepository) GetById(ctx context.Context, id string) (*domain.Habit, error) {

	getById := `Select * FROM habits WHERE id = $1`

	row := h.dbConn.QueryRow(ctx, getById, id)

	var habit domain.Habit
	err := row.Scan(&habit.Id, &habit.Name, &habit.Description, &habit.UserId)
	if err != nil {
		fmt.Printf("habitRepository.GetById error: %s\n", err)
		return nil, err
	}

	return &habit, nil
}

func (h *habitRepository) Save(ctx context.Context, habit *domain.Habit) error {

	insertSql := `INSERT INTO habits (name, description, user_id) VALUES ($1, $2, $3)`

	_, err := h.dbConn.Exec(ctx, insertSql, habit.Name, habit.Description, habit.UserId)
	if err != nil {
		fmt.Printf("habitRepository.Save error: %s\n", err)
		return err
	}

	return nil
}

func (h *habitRepository) Update(ctx context.Context, habit *domain.Habit) error {

	_, err := h.GetById(ctx, habit.Id)
	if err != nil {
		return err
	}

	updateSql := `UPDATE habits SET name = $1, description = $2 WHERE id = $3`

	_, err = h.dbConn.Exec(ctx, updateSql, habit.Name, habit.Description, habit.Id)
	if err != nil {
		fmt.Printf("habitRepository.Update error: %s\n", err)
		return err
	}

	return nil
}

func (h *habitRepository) Delete(ctx context.Context, id string) error {

	_, err := h.GetById(ctx, id)
	if err != nil {
		return err
	}

	deleteSql := `DELETE FROM habits WHERE id = $1`

	_, err = h.dbConn.Exec(ctx, deleteSql, id)
	if err != nil {
		fmt.Printf("habitRepository.Delete error: %s\n", err)
		return err
	}

	return nil
}
