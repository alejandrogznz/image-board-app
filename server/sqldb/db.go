// With guidance from 
// https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/
package sqldb

import (
	"os"
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
)
// DB is a global variable to hold db connection

func dsn() string {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	hostname := os.Getenv("HOST")
	dbname	 := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func ConnectDB() *sql.DB {
	err := godotenv.Load("secrets.env")
	if err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("mysql", dsn())
	if err != nil {
		panic(err.Error())
	}
	return db
}
