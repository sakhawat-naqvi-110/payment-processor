package database

import (
	"database/sql"
	"gorm.io/gorm"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *gorm.DB
}

func New() *Database {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("DB open error", "msg", err)
	}

	return &Database{
		db: db,
	}
}

func (d *Database) CreateTable() error {
	// Start a DB transaction
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(`CREATE TABLE Person(
		Id int not null,
		Name varchar not null,
		DateOfBirth date not null,
		Gender bit not null,
		PRIMARY KEY( Id )
	  );`)
	if err != nil {
		return err
	}

	log.Println(res)

	// Commit the transaction
	return tx.Commit()
}

func (d *Database) DropTable() error {
	// Start a DB transaction
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(`DROP TABLE Person;`)
	if err != nil {
		return err
	}

	log.Println(res)

	// Commit the transaction
	return tx.Commit()
}
