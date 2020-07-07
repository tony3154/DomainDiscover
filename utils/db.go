package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open("mysql", "go:Asong123!@#@tcp(192.168.182.200:3306)/domain")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("数据库连接成功")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS user(id int(10) primary key auto_increment,name varchar(20),password varchar(30));")
	if err != nil {
		fmt.Println("create table user failed:", err.Error())
		return
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS domain(id int(10) primary key auto_increment,domainName varchar(50),project varchar(50),service varchar(50),CDN varchar(50),HTTPS varchar(50),backend varchar(50),whiteList varchar(50),whiteListLocation varchar(50),notes varchar(200));")
	if err != nil {
		fmt.Println("create table domain failed:", err.Error())
		return
	}
}
