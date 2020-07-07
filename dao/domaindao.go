package dao

import (
	"domain/utils"
	"fmt"
	"strconv"

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
	sqlStr := "select id,domainName,project,service,CDN,HTTPS,backend,whiteList,whiteListLocation,notes from domain where 1=1 "
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		fmt.Println("select db failed,err:", err)
		return
	}

	//封装查询返回
	dbdomains := DBreturn(rows)
	table.Output(dbdomains)
}

//根据域名查询
func SelectDomain(name string) {
	// dbdomain := &model.Domain{}
	sqlStr := "select id,domainName,project,service,CDN,HTTPS,backend,whiteList,whiteListLocation,notes from domain where domainName=? "
	rows, err := utils.DB.Query(sqlStr, name)
	if err != nil {
		fmt.Println("select db failed,err:", err)
		return
	}

	//封装查询返回
	dbdomains := DBreturn(rows)
	table.Output(dbdomains)

}

//根据项目
func SelectProject(project string) {
	sqlStr := "select id,domainName,project,service,CDN,HTTPS,backend,whiteList,whiteListLocation,notes from domain where project=? "
	rows, err := utils.DB.Query(sqlStr, project)
	if err != nil {
		fmt.Println("select db failed,err:", err)
		return
	}

	//封装查询返回
	dbdomains := DBreturn(rows)
	table.Output(dbdomains)
}

//添加域名到数据库
func AddDomain() {

	//对输入赋值到结构体
	newdomain := InputDomain()

	//对比数据库域名是否已经存在
	id := JudgmentDomainName(newdomain.DomainName)

	if id != 0 {
		fmt.Println("域名已存在")
		return
	}

	//域名信息存入数据库
	fmt.Println("域名信息存入数据库")
	sqlStr := "insert into domain(domainName,project,service,CDN,HTTPS,backend,whiteList,whiteListLocation,notes) values(?,?,?,?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, newdomain.DomainName, newdomain.Project, newdomain.Service, newdomain.CDN, newdomain.HTTPS, newdomain.Backend, newdomain.WhiteList, newdomain.WhiteListLocation, newdomain.Notes)
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

//修改域名菜单项
func ModifyMenu() {
	//菜单，以ID或域名进行修改

	for {
		input := Inputstr("请输入查询方式：ID输入I/域名输入Y，Q返回")
		switch {
		case input == "I" || input == "i":
			inputID := Inputstr("请输入ID：")
			ID, err := strconv.Atoi(inputID)
			if err != nil {
				fmt.Println("输入错误，请重试")
				continue
			}
			ModifyID(ID)
			return
		case input == "Y" || input == "y":
			inputDomain := Inputstr("请输入域名：")
			ModifyDomainName(inputDomain)
			return
		case input == "Q" || input == "q":
			return
		default:
			fmt.Println("输入错误，请重新输入\n\n\n")
		}
	}

}

//按ID修改
func ModifyID(id int) {

	dbid := JudgmentID(id)
	if dbid == 0 {
		fmt.Println("输入的ID不存在")
		return
	}

	ModifyDomain(dbid)

}

func ModifyDomainName(domain string) {
	id := JudgmentDomainName(domain)
	if id == 0 {
		fmt.Println("域名不存在")
		return
	}

	ModifyDomain(id)
}

func ImportDomain() {
	domains := ReadExcel()
	for _, domain := range domains {

		id := JudgmentDomainName(domain.DomainName)
		if id != 0 {
			fmt.Println(domain.DomainName, "已经存在,未修改")
			continue
		}

		sqlStr := "insert into domain(domainName,project,service,CDN,HTTPS,backend,whiteList,whiteListLocation,notes) values(?,?,?,?,?,?,?,?,?)"
		_, err := utils.DB.Exec(sqlStr, domain.DomainName, domain.Project, domain.Service, domain.CDN, domain.HTTPS, domain.Backend, domain.WhiteList, domain.WhiteListLocation, domain.Notes)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(domain.DomainName, "存入数据库")
		}

	}
}
