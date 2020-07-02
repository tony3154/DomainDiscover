package main

import (
	. "domain/dao"
	_ "domain/utils"
	"fmt"
	"os"
)

func main() {

	//主菜单变量
	var menu string
	menu = `
欢迎使用域名查询系统
1.注册
2.登录
3.修改
4.退出
`

	for {
		fmt.Println(menu)
		var inputNumber int
		fmt.Scan(&inputNumber)

		switch inputNumber {
		case 1:
			fmt.Println("注册用户")
			Regist()
		case 2:
			fmt.Println("登录")
			Login()
		case 3:
			fmt.Println("修改密码")
		case 4:
			fmt.Println("欢迎使用，下次再见！！")
			os.Exit(0)
		default:
			fmt.Println("输入无效，请重新输入")
		}

	}

}
