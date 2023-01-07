package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type database struct {
	DB     *sql.DB
	err    error
	errMsg string
}

func (db *database) connectDatabase() {
	fmt.Println("address database server:", os.Getenv("DATABASE_URL"))
	db.DB, db.err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if db.err != nil {
		log.Fatal("Connect to database error", db.err)
	}
}

func (db *database) createDatabase() {
	createTB := `CREATE TABLE IF NOT EXISTS expenses ( id SERIAL PRIMARY KEY, title TEXT, amount FLOAT, note TEXT, tags TEXT[] )`
	_, db.err = db.DB.Exec(createTB)
	if db.err != nil {
		db.errMsg = db.err.Error()
		log.Fatal("cant`t create table", db.err)
	}
	log.Println("Okey Database it Have Table")
}

func (db *database) InitDatabase() {
	db.connectDatabase()
	db.createDatabase()
}

func (db *database) CloseDatabase() {
	db.DB.Close()
}
