package dao

import (
	"domain/model"
	"domain/utils"
	"fmt"

	"github.com/modood/table"
)

//查询菜单，分三个，列出所有，以域名查询，以项目名查询
func SelectMenu() {
LOOP:
	for {
		input := Inputstr("根据域名查询请输入Y，根据项目查询请输入P，显示所有请输入A,退出请输入Q")
		switch {
		case input == "Y" || input == "y":
			domainName := Inputstr("请输入域名")
			SelectDomain(domainName)
		case input == "P" || input == "p":
			project := Inputstr("请输入项目")
			SelectProject(project)
		case input == "A" || input == "a":
			SelectAll()
		case input == "Q" || input == "q":
			break LOOP
		default:
			fmt.Println("输入错误，请重试")
		}
	}
}

//显示所有域名
func SelectAll() {
	sqlStr := "select id,domainName,project,service,CDN,sslState,whiteList,whiteListLocation from domain where 1=1 "
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		fmt.Println("select db failed,err:", err)
		return
	}

	//封装查询返回
	dbdomains := DBreturn(rows)
	t := table.Table(dbdomains)
	fmt.Println(t)
}

//根据域名查询
func SelectDomain(name string) {
	// dbdomain := &model.Domain{}
	sqlStr := "select id,domainName,project,service,CDN,sslState,whiteList,whiteListLocation from domain where domainName=? "
	rows, err := utils.DB.Query(sqlStr, name)
	if err != nil {
		fmt.Println("select db failed,err:", err)
		return
	}

	//封装查询返回
	dbdomains := DBreturn(rows)

	t := table.Table(dbdomains)
	fmt.Println(t)
}

//根据项目
func SelectProject(project string) {
	sqlStr := "select id,domainName,project,service,CDN,sslState,whiteList,whiteListLocation from domain where project=? "
	rows, err := utils.DB.Query(sqlStr, project)
	if err != nil {
		fmt.Println("select db failed,err:", err)
		return
	}

	//封装查询返回
	dbdomains := DBreturn(rows)
	t := table.Table(dbdomains)
	fmt.Println(t)
}

//添加域名到数据库
func AddDomain() {

	//对输入进行赋值到结构体
	newdomain := &model.Domain{}
	newdomain.DomainName = Inputstr("请输入域名")
	newdomain.Project = Inputstr("请输入所属项目")
	newdomain.Service = Inputstr("请输入所属服务")
	newdomain.CDN = Inputstr("请输入CDN厂商")

LOOP:
	for {
		var sslState string
		sslState = Inputstr("是否HTTPS（Y/N）")

		switch {
		case sslState == "Y" || sslState == "y":
			newdomain.SslState = true
			break LOOP
		case sslState == "N" || sslState == "n":
			newdomain.SslState = false
			break LOOP
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}

	newdomain.WhiteList = Inputstr("加白位置(SLB/安全组/iptables)")
	newdomain.WhiteListLocation = Inputstr("加白位置IP")

	//对比数据库域名是否已经存在
	id := JudgmentDomainName(newdomain.DomainName)

	if id != 0 {
		fmt.Println("域名已存在")
		return
	}

	//域名信息存入数据库
	fmt.Println("域名信息存入数据库")
	sqlStr := "insert into domain(domainName,project,service,CDN,sslState,whiteList,whiteListLocation) values(?,?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, newdomain.DomainName, newdomain.Project, newdomain.Service, newdomain.CDN, newdomain.SslState, newdomain.WhiteList, newdomain.WhiteListLocation)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteDomain() {

	name := Inputstr("请输入域名")
	//对比数据库域名是否已经存在
	id := JudgmentDomainName(name)

	if id == 0 {
		fmt.Println("域名不存在")
		return
	}

	//根据返回的ID删除记录
	strsql := "DELETE FROM domain WHERE id=?"
	_, err := utils.DB.Exec(strsql, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, "删除成功")

}

func ModifyDomain() {
	//

}
