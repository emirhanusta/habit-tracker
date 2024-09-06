package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"habit-tracker/internal/domain"
)

type IReminderRepository interface {
	GetAllByHabitId(ctx context.Context, habitId string) ([]domain.Reminder, error)
	GetById(ctx context.Context, id string) (*domain.Reminder, error)
	Save(ctx context.Context, reminder *domain.Reminder) error
	Update(ctx context.Context, reminder *domain.Reminder) error
	Delete(ctx context.Context, id string) error
}

type reminderRepository struct {
	dbConn *pgx.Conn
}

func NewReminderRepository(dbConn *pgx.Conn) IReminderRepository {
	return &reminderRepository{
		dbConn: dbConn,
	}
}

func (r *reminderRepository) GetAllByHabitId(ctx context.Context, habitId string) ([]domain.Reminder, error) {

	getAllByHabitId := `SELECT * FROM reminders WHERE habit_id = $1`

	rows, err := r.dbConn.Query(ctx, getAllByHabitId, habitId)
	if err != nil {
		fmt.Printf("reminderRepository.GetAllByHabitId error: %s\n", err)
		return nil, err
	}

	var reminders []domain.Reminder

	for rows.Next() {
		var reminder domain.Reminder
		err = rows.Scan(&reminder.Id, &reminder.HabitId, &reminder.RemindAt, &reminder.Message)
		if err != nil {
			fmt.Printf("reminderRepository.GetAllByHabitId error: %s\n", err)
			return nil, err
		}
		reminders = append(reminders, reminder)
	}

	return reminders, nil
}

func (r *reminderRepository) GetById(ctx context.Context, id string) (*domain.Reminder, error) {

	getById := `SELECT * FROM reminders WHERE id = $1`

	row := r.dbConn.QueryRow(ctx, getById, id)

	var reminder domain.Reminder
	err := row.Scan(&reminder.Id, &reminder.HabitId, &reminder.RemindAt, &reminder.Message)
	if err != nil {
		fmt.Printf("reminderRepository.GetById error: %s\n", err)
		return nil, err
	}

	return &reminder, nil
}

func (r *reminderRepository) Save(ctx context.Context, reminder *domain.Reminder) error {

	insertSql := `INSERT INTO reminders (habit_id, remind_at, message) VALUES ($1, $2, $3)`

	_, err := r.dbConn.Exec(ctx, insertSql, reminder.HabitId, reminder.RemindAt, reminder.Message)
	if err != nil {
		fmt.Printf("reminderRepository.Save error: %s\n", err)
		return err
	}

	return nil
}

func (r *reminderRepository) Update(ctx context.Context, reminder *domain.Reminder) error {

	_, err := r.GetById(ctx, reminder.Id)
	if err != nil {
		return err
	}

	updateSql := `UPDATE reminders SET remind_at = $1, message = $2 WHERE id = $3`

	_, err = r.dbConn.Exec(ctx, updateSql, reminder.RemindAt, reminder.Message, reminder.Id)
	if err != nil {
		fmt.Printf("reminderRepository.Update error: %s\n", err)
		return err
	}

	return nil
}

func (r *reminderRepository) Delete(ctx context.Context, id string) error {

	_, err := r.GetById(ctx, id)
	if err != nil {
		return err
	}

	deleteSql := `DELETE FROM reminders WHERE id = $1`

	_, err = r.dbConn.Exec(ctx, deleteSql, id)
	if err != nil {
		fmt.Printf("reminderRepository.Delete error: %s\n", err)
		return err
	}

	return nil
}
