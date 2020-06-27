package dao

import (
	"fmt"
	"os"
)

func DomainMain() {
	var menu string
	menu=`
1.查询域名
2.增加域名
3.删除域名
4.修改域名
5.退出
`


	for {
		fmt.Println(menu)
		var inputNumber int
		fmt.Scan(&inputNumber)

		switch inputNumber {
		case 1:
			fmt.Println("查询域名")

		case 2:
			fmt.Println("增加域名")
		case 3:
			fmt.Println("删除域名")
		case 4:
			fmt.Println("修改域名")
		case 5:
			fmt.Println("欢迎使用，下次再见！！")
			os.Exit(0)
		default:
			fmt.Println("输入无效，请重新输入")
		}
	}


}