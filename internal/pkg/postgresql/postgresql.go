package postgresql

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
)

func ConnectPostgres(username, password, host, port, dbName string) (*pgx.Conn, error) {
	// Connect to PostgreSQL
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
