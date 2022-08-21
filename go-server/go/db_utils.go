package swagger

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {

	host, port, user, password, dbname := getEnvVars()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func closeDB(db *sql.DB) {
	db.Close()
}
