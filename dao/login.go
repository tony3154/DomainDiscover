package dao

import (
	"fmt"
)

func Login() {
	var name,password string
	fmt.Println("请输入账号：")
	fmt.Scan(&name)
	fmt.Println("请输入密码：")
	fmt.Scan(&password)

	_, err := CheckUserNameAndPassword(name,password)
	if err!=nil {
		fmt.Println("登录失败",err)

	}else {
		fmt.Println("登录成功")
		DomainMain()
	}
}
