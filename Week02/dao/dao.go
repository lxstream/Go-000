package dao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int64
	Name string
}

var (
	db  *sql.DB
	dsn string
)

func init() {
	dsn = "root:123!@#Abc@tcp(192.168.247.131)/testdb?charset=utf8mb4&parseTime=True"
}

func SelectUser(id int) (*User, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	sqlStr := "select id,name from user where id=?"
	var u User

	err = db.QueryRow(sqlStr, id).Scan(&u.Id, &u.Name)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
