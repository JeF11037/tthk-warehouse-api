package modules

import (
	"database/sql"
	"fmt"
	"main/models"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func Connect() {
	database, err := sql.Open("mysql", "root:@/tthk_warehouse")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
	DB = database
}

func VerifyUser(user models.User) (bool, error) {
	rows, err := DB.Query("SELECT * FROM user WHERE username = ? and password = ?", user.Username, user.Password)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		rows.Close()
		return true, err
	}
	return false, err
}

func CreateUser(user models.User) error {
	_, err := DB.Exec("INSERT INTO user(username, password, first_name, last_name, address_id, payment_id) VALUES(?, ?, ?, ?, ?, ?)",
		user.Username, user.Password, user.FirstName, user.LastName, user.Address, user.Payment)
	if err != nil {
		return err
	}
	return err
}
