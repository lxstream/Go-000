package dao

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
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
	dsn = "root:123123@tcp(127.0.0.1)/testdb?charset=utf8mb4&parseTime=True"
}

func SelectUser(id int) (*User, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	sqlStr := "select id,name from user where id=?"
	var u User

	err = db.QueryRow(sqlStr, id).Scan(&u.Id, &u.Name)
	// if err != nil {
	// 	return nil, err
	// }
	// return &u, nil

	//
	switch {
	case err == sql.ErrNoRows:
		errMsg := "Can not find user with id:" + strconv.Itoa(id)
		return nil, errors.Wrap(err, errMsg)
	case err != nil:
		errMsg := "something wrong in db query!"
		return nil, errors.Wrap(err, errMsg)
	default:
		return &u, nil
	}
}
