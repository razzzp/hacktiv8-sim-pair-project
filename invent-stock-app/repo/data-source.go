package repo

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateDBInstance() (*sql.DB, error) {
	//db connection
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// open connection
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	// ping DB to check connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
