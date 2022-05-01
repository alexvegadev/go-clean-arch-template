package database

import (
	"database/sql"
	"fmt"
	"go-clean-arch/internal/platform/log"

	_ "github.com/go-sql-driver/mysql"
)

func New(user string, password string, host string, port string, database string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?allowNativePasswords=true&parseTime=true&utcLoc", user, password,
		host, port, database)

	log.Infof("Connecting to database using user %s and host %s", user, host)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	log.Info("Sending ping to database and waiting for response...")

	check := db.Ping()

	if check != nil {
		return nil, check
	}

	return db, err
}
