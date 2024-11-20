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

func (db *DB) GetRecordsByAddress(ctx context.Context, address string) ([]model.Attendance, error) {
	rows, err := db.pool.Query(ctx,
		"SELECT id, address, created_at FROM attendance WHERE address ILIKE $1",
		address)
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

func (db *DB) GetUniqueAddressesByDay(ctx context.Context, date string) (int, []string, error) {
	rows, err := db.pool.Query(ctx,
		`SELECT COUNT(DISTINCT address) as unique_address_count, 
			ARRAY_AGG(DISTINCT address) as unique_addresses 
		 FROM attendance 
		 WHERE DATE(created_at) = $1`,
		date)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	var uniqueCount int
	var uniqueAddresses []string

	if rows.Next() {
		if err := rows.Scan(&uniqueCount, &uniqueAddresses); err != nil {
			return 0, nil, err
		}
	}

	return uniqueCount, uniqueAddresses, nil
}
