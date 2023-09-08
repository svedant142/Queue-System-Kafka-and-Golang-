package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DbClient *sql.DB

func InitMYSQL() error {
	var err error
	DbClient, err = sql.Open("mysql", "root:{:Password}@tcp(127.0.0.1:3306)/DBname")
	if err != nil {
		fmt.Println(err)
		return err	
	}
	return nil
}