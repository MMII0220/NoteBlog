package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// var db *sqlx.DB

func StartDBConnection() (*sqlx.DB, error) {
	// var err error
	connStr := "host=localhost port=5432 user=postgres dbname=test password=1234 sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error in connecting to db: %v", err)
	}
	return db, nil
}

func CloseDB(db *sqlx.DB) error {
	if db != nil {
		err := db.Close()
		if err != nil {
			return fmt.Errorf("error in closing db %v: ", err)
		}
	}
	return nil
}

// func GetDBConnection() *sqlx.DB {
// 	return db
// }
