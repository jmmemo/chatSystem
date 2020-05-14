package main

import (
	"chatroom/common"
	process2 "chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

////////
func (this *Processor) serverProcessMes(msg *common.Message) (err error) {

	switch msg.Type {
	case common.LoginMesType:
		//登录
		up := process2.UserProcess{
			Conn: this.Conn,
		}
		err := up.ServerProcessLogin(msg)

		if err != nil {
			fmt.Println("serverProcessLogin(conn, msg),err=", err)
			///////////return
		}
	case common.RegisterMesType:
		//注册

	default:
		fmt.Println("无法处理")
	}
	return
}

func (this *Processor) process3() (err error) {
	fmt.Println("等待读取client发送的数据...")
	//tf := &utils.Transfer{
	//	Conn: this.Conn,
	//}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	msg, err := tf.ReadPkg()

	if err == io.EOF {
		fmt.Println("客户端退出了..")
		return
	}
	if err != nil {
		fmt.Println("readPkg(conn),err=", err)
		return
	}
	///////////////////////

	err = this.serverProcessMes(&msg)
	if err != nil {
		fmt.Println("serverProcessMes(conn, &msg),err", err)
		return
	}
	fmt.Println("我在process里拿到里msg", msg)
	return
}
