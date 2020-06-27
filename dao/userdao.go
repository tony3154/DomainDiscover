package dao

import (
	"domain/model"
	"domain/utils"
	"fmt"
)

//向数据库核对账号密码
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	//sql语句
	fmt.Println("核对账号密码")
	sqlStr := "select id,name,password from user where name=? and password=?"
	row := utils.DB.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, )
	fmt.Println(user.ID)
	err:=error(nil)
	if user.ID==0 {
		return user,fmt.Errorf("账号密码核对错误")
	}
	return user, err
}

//查询user是否存在
func CheckUserName(username string) (*model.User, error) {
	//sql语句
	fmt.Println("核对账号")
	sqlStr := "select id,name,password from user where username=? "
	row := utils.DB.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password)
	fmt.Println(user.ID,user.Username,user.Password)
	fmt.Println(user.ID)
	return user, nil
}


//保存user,向数据库中写入用户信息
func SaveUser(username string, password string) error {
	//sql语句
	sqlStr := "insert into user(name,password) values(?,?)"
	//执行
	_, err := utils.DB.Exec(sqlStr, username, password)
	if err != nil {
		return err
	}
	return nil
}