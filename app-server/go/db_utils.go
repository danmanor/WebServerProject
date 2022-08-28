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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (userId UUID DEFAULT gen_random_uuid (), name VARCHAR (50), email VARCHAR (255) UNIQUE, PRIMARY KEY(userId))")

	return db
}

func closeDB(db *sql.DB) {
	db.Close()
}
