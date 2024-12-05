package db

import (
	"context"
	"fmt"
	"github.com/JollyGrin/postgres-attendance/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
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

// NOTE:  old insert, now has duplication check
// @deprecated
// func (db *DB) RecordAttendance(ctx context.Context, address string, location string, metaverse model.MetaverseType, entranceStatus model.EntranceStatusType) error {
// 	_, err := db.pool.Exec(ctx,
// 		"INSERT INTO attendance (address, created_at, location, metaverse, entrance_status) VALUES ($1, NOW(), $2, $3, $4)",
// 		address, location, metaverse, entranceStatus)
// 	if err != nil {
// 		return fmt.Errorf("failed to record attendance: %w", err)
// 	}
// 	return nil
// }

func (db *DB) RecordAttendance(ctx context.Context, address string, location string, metaverse model.MetaverseType, entranceStatus model.EntranceStatusType) (bool, error) {
	// Log the values being recorded
	logrus.WithFields(logrus.Fields{
		"address $1":        address,
		"location $2":       location,
		"metaverse $3":      metaverse,
		"entranceStatus $4": entranceStatus,
	}).Info("Attempting to record attendance")

	// Use a query to insert the record only if a similar one does not exist within the last 5 seconds
	commandTag, err := db.pool.Exec(ctx,
		`INSERT INTO attendance (address, created_at, location, metaverse, entrance_status)
		SELECT $1, NOW(), $2, $3, $4
		WHERE NOT EXISTS (
			SELECT 1
			FROM attendance
			WHERE address = $1
			  AND location = $2
        AND metaverse = $3::VARCHAR
        AND entrance_status = $4::VARCHAR
			  AND created_at > NOW() - INTERVAL '5 seconds'
		)`,
		address, location, metaverse, entranceStatus)

	if err != nil {
		return false, fmt.Errorf("failed to record attendance: %w", err)
	}

	// Check if a row was inserted
	// Determine if the record was inserted
	return commandTag.RowsAffected() > 0, nil
}

type UserDuration struct {
	Address   string    `json:"address"`
	EnterTime time.Time `json:"enter_time"`
	ExitTime  time.Time `json:"exit_time"`
	Duration  float64   `json:"duration"`
}

func (db *DB) GetUserDurationsByDay(ctx context.Context, date string) ([]UserDuration, error) {
	// Query to get all enter/exit events for the day
	rows, err := db.pool.Query(ctx,
		`WITH events AS (
            SELECT address, created_at, entrance_status,
                   ROW_NUMBER() OVER (PARTITION BY address ORDER BY created_at) as rn
            FROM attendance
            WHERE DATE(created_at) = $1
        ),
        paired_events AS (
            SELECT 
                e1.address, 
                e1.created_at AS enter_time, 
                COALESCE(e2.created_at, e1.created_at + INTERVAL '10 minutes') AS exit_time
            FROM events e1
            LEFT JOIN events e2 ON e1.address = e2.address AND e1.rn = e2.rn - 1 AND e2.entrance_status = 'EXIT'
            WHERE e1.entrance_status = 'ENTER'
        )
        SELECT address, enter_time, exit_time, exit_time - enter_time AS duration
        FROM paired_events`,
		date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userDurations []UserDuration

	for rows.Next() {
		var ud UserDuration
		var duration time.Duration // New variable to hold the duration from the database
		if err := rows.Scan(&ud.Address, &ud.EnterTime, &ud.ExitTime, &duration); err != nil {
			return nil, err
		}
		ud.Duration = duration.Seconds() // Convert to seconds before appending
		userDurations = append(userDurations, ud)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userDurations, nil
}
