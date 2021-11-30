package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	db *sql.DB
)

func Connect() {
	var err error
	cfg := mysql.Config{
		User:   "root@localhost",
		Passwd: "",
		Net:    "tcp",
		Addr:   "localhost/phpmyadmin/",
		DBName: "tthk_warehouse",
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	if err != nil {
		panic(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
