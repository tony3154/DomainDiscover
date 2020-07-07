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

//根据ID判断域名是否存在
func JudgmentID(id int) int {
	var dbid int
	sqlStr := "select id from domain where id=? "
	row := utils.DB.QueryRow(sqlStr, id)
	row.Scan(&dbid)
	return dbid
}

//域名查询返回
func DBreturn(rows *sql.Rows) []*model.Domain {
	dbdomains := []*model.Domain{}

	for rows.Next() {
		dbdomain := &model.Domain{}
		err := rows.Scan(&dbdomain.ID, &dbdomain.DomainName, &dbdomain.Project, &dbdomain.Service, &dbdomain.CDN, &dbdomain.HTTPS, &dbdomain.Backend, &dbdomain.WhiteList, &dbdomain.WhiteListLocation, &dbdomain.Notes)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(dbdomain.ID, dbdomain.DomainName, dbdomain.Project, dbdomain.Service, dbdomain.CDN, dbdomain.HTTPS, dbdomain.WhiteList, dbdomain.WhiteListLocation)

		dbdomains = append(dbdomains, dbdomain)
	}
	return dbdomains
}

func InputDomain() *model.Domain {
	//对输入赋值到结构体
	newdomain := &model.Domain{}
	newdomain.DomainName = Inputstr("请输入域名")
	newdomain.Project = Inputstr("请输入所属项目")
	newdomain.Service = Inputstr("请输入所属服务")
	newdomain.CDN = Inputstr("请输入CDN厂商")
	newdomain.Backend = Inputstr("请输入回源IP")
LOOP:
	for {
		var HTTPS string
		HTTPS = Inputstr("有无HTTPS（有/无）")

		switch {
		case HTTPS == "有" || HTTPS == "无":
			newdomain.HTTPS = HTTPS
			break LOOP
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}

	newdomain.WhiteList = Inputstr("加白位置(SLB/安全组/iptables)")
	newdomain.WhiteListLocation = Inputstr("加白位置IP")
	newdomain.Notes = Inputstr("请输入备注信息")
	return newdomain
}

func ModifyDomain(id int) {
	//获取输入的域名信息
	fmt.Println("请输入修改后的信息")
	newdomain := InputDomain()

	//对比修改后的域名在数据库是否已经存在,存在则不允许修改
	newid := JudgmentDomainName(newdomain.DomainName)
	if newid != 0 {
		fmt.Println("域名已存在")
		return
	}

	sqlStr := "update domain set domainName=?,project=?,service=?,CDN=?,HTTPS=?,backend=?,whiteList=?,whiteListLocation=?,notes=? where id=?"
	_, err := utils.DB.Exec(sqlStr, newdomain.DomainName, newdomain.Project, newdomain.Service, newdomain.CDN, newdomain.HTTPS, newdomain.Backend, newdomain.WhiteList, newdomain.WhiteListLocation, newdomain.Notes, id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n\n\n")
	fmt.Print("修改成功\n\n\n\n")
}
