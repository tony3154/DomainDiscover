package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open("mysql", "domain:3154@tcp(192.168.0.100:3306)/domain")
	fmt.Println("数据库连接成功")
	if err != nil {
		fmt.Println(err)
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS user(id int(10) primary key auto_increment,name varchar(20),password varchar(30));")
	if err != nil {
		fmt.Println("create table user failed:", err.Error())
		return
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS domain(id int(10) primary key auto_increment,domainName varchar(50),project varchar(50),service varchar(50),CDN varchar(50),sslState bool,whiteList varchar(50),whiteListLocation varchar(50));")
	if err != nil {
		fmt.Println("create table domain failed:", err.Error())
		return
	}


}