package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
	"os"
)

type user struct {
	id         int
	name       string
	departname string
	created    string
}

var u user
var db *sql.DB

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return nil
}

func main() {
	//初始化数据库连接
	initDB()
	defer db.Close()

	id := 0
	//查询
	u, err := query(id)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(0)
	}

	fmt.Printf("id:%d,name:%s,departname:%s,created:%s", u.id, u.name, u.departname, u.created)

}

func query(id int) (*user, error) {
	err := db.QueryRow("select * from user_info where uid=?", id).Scan(&u.id, &u.name, &u.departname, &u.created)
	if err != nil {
		return nil, xerrors.Wrapf(err, fmt.Sprintf("id:%d 查询错误", id))
	}
	return &u, nil
}
