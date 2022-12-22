package main

import (
	"fmt"
	"go_code/project01/main/tongxunxitong/client/process"
	"os"
)

func main() {

	//接收用户的选择
	var key int
	//判断是否还要继续显示菜单
	//var loop = true
	//定义用户账号密码
	var userId int
	var userPwd string
	var userName string

	for true {
		fmt.Println("----------欢迎使用海量用户聊天系统----------")
		fmt.Println("\t\t1.进入聊天")
		fmt.Println("\t\t2.注册用户")
		fmt.Println("\t\t3.退出系统")
		fmt.Println("请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登入聊天")
			fmt.Println("请输入您的账号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入您的密码")
			fmt.Scanf("%s\n", &userPwd)
			//创建一个UserProcess实例
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("up.Login err=", err)
			}
			//loop = false

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入您的账号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入您的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入您的用户名")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
			//loop = false

		default:
			fmt.Println("您的输入有误，请重新输入")

		}
	}
	//根据用户的选择，显示新的提示信息
	/*if key == 1 {

		fmt.Println("请输入您的账号")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入您的密码")
		fmt.Scanf("%s\n", &userPwd)
		//登录界面写到另外一个包
		err := Login(userId, userPwd)
		if err != nil {
			fmt.Printf("登录失败，err=%v", err)
		} else {
			fmt.Println("\n登陆成功")
		}
	} else if key == 2 {
		fmt.Println("请注册新账户")
	}*/
}
