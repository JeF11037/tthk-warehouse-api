package modules

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Connect() {
	var err error
	DB, err = gorm.Open("mysql", "root:@/tthk_warehouse")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
}

func Disconnect() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database disconnected")
}
