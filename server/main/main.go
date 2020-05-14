package main

import (
	"chatroom/server/model"
	"fmt"
	"io"
	"net"
	"time"
)

////登录
//func serverProcessLogin(conn net.Conn, msg *common.Message) (err error) {
//	var loginMes common.LoginMes
//	err = json.Unmarshal([]byte(msg.Data), &loginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal([]byte(msg.Data), loginMes),err=", err)
//		return
//	}
//	//构建返回消息
//	var resMes common.Message
//	resMes.Type = common.LoginResMesType
//
//	var loginResMes common.LoginResMes
//
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//		//合法
//		loginResMes.Code = 200
//	} else {
//		//不合法
//		loginResMes.Code = 500
//		loginResMes.Error = "用户不存在，请注册再使用"
//	}
//	//为返回resMes做准备
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.Marshal(loginResMes),err=", err)
//		return
//	}
//	resMes.Data = string(data)
//	//序列化准备发送
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json.Marshal(resMes),err=", err)
//		return
//	}
//	/////writePkg
//	err = utils.WritePkg(conn, data)
//	if err != nil {
//		fmt.Println("writePkg(conn, data),err=", err)
//		return
//	}
//	return
//}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		pro := &Processor{
			Conn: conn,
		}
		err := pro.process3()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("客户与服务端通讯协程错误,err=", err)
			return
		}
	}
}

//初始化Dao对象
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	//初始化
	initPool("localhost:6379", 16, 0, time.Second*300)
	////////
	initUserDao()
	fmt.Println("服务器监听8889。。")
	listen, err := net.Listen("tcp", "localhost:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen(8889),err=", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept(),err=", err)
			return
		}
		go process(conn)
		////////
	}

}
