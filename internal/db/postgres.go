package db

import (
	"context"
	"time"

	"github.com/JollyGrin/postgres-attendance/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(dbURL string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return &DB{pool: pool}, nil
}

func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) GetTodayAttendance(ctx context.Context) ([]model.Attendance, error) {
	today := time.Now().Format("2006-01-02")

	rows, err := db.pool.Query(ctx,
		"SELECT id, address, created_at FROM attendance WHERE DATE(created_at) = $1",
		today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []model.Attendance
	for rows.Next() {
		var record model.Attendance
		if err := rows.Scan(&record.ID, &record.Address, &record.Created_At); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
