package main

import (
	"chatroom/client/process"
	"fmt"
	"os"
)

var (
	key     int
	userId  int
	userPwd string
)

func main() {
	fmt.Println("·············欢迎登录多人运动系统··············")
	for {

		fmt.Println("\t\t\t1登录聊天系统")
		fmt.Println("\t\t\t2注册用户")
		fmt.Println("\t\t\t3退出系统")
		fmt.Println("请选择1-3:")
		fmt.Println("··························")
		_, _ = fmt.Scanln(&key)

		switch key {
		case 1:
			//从终端获取id和密码输入
			fmt.Println("请输入用户id:")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanln(&userPwd)
			fmt.Printf("id:%d\t密码:%s\n", userId, userPwd)

			//
			//err := login(userId, userPwd)
			//if err != nil {
			//	fmt.Println("login(userId, userPwd),err=", err)
			//	return
			//}

			//up := &process.UserProcess{
			//
			//}
			up:=&process.UserProcess{}

			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("Login(userId, userPwd),err=", err)
			}
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出")
			os.Exit(1)
		default:
			fmt.Println("输入出错")
		}
	}
}
