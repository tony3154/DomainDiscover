package dao

import "fmt"

func Regist(){

	var name,password string
	fmt.Println("请输入账号：")
	fmt.Scan(&name)
	fmt.Println("请输入密码：")
	fmt.Scan(&password)

	//_,err := CheckUserName(name)
	//if err!=nil{
	//	fmt.Println("用户已存在，请更换用户名")
	//	return
	//}
	err := SaveUser(name,password)
	if err!= nil {
		fmt.Println(err)
	}else{
		fmt.Println("用户注册成功")
	}




}