package dao

import (
	"database/sql"
	"domain/model"
	"domain/utils"
	"fmt"
)

//输入赋值，
func Inputstr(str string) string {
	fmt.Println(str)
	var input string
	fmt.Scan(&input)
	return input
}

//判断域名是否在数据库存在,输入域名，返回域名的ID，为0则不存在
func JudgmentDomainName(name string) int {
	var id int
	sqlStr := "select id from domain where domainName=? "
	row := utils.DB.QueryRow(sqlStr, name)
	row.Scan(&id)
	return id
}

//域名查询返回
func DBreturn(rows *sql.Rows) []*model.Domain {
	dbdomains := []*model.Domain{}

	for rows.Next() {
		dbdomain := &model.Domain{}
		err := rows.Scan(&dbdomain.ID, &dbdomain.DomainName, &dbdomain.Project, &dbdomain.Service, &dbdomain.CDN, &dbdomain.SslState, &dbdomain.WhiteList, &dbdomain.WhiteListLocation)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(dbdomain.ID, dbdomain.DomainName, dbdomain.Project, dbdomain.Service, dbdomain.CDN, dbdomain.SslState, dbdomain.WhiteList, dbdomain.WhiteListLocation)

		dbdomains = append(dbdomains, dbdomain)
	}
	return dbdomains
}
