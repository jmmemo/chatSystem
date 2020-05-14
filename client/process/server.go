package process

import (
	"chatroom/client/utils"
	"fmt"
	"net"
	"os"
)

func showMenu() {
	fmt.Println("-----------恭喜xxx登录成功")
	fmt.Println("-----------1显示在线用户列表")
	fmt.Println("-----------2发送消息")
	fmt.Println("-----------3信息列表")
	fmt.Println("-----------4退出")
	fmt.Println("请选择1-4:")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统...")
		os.Exit(1)
	default:
		fmt.Println("你输入的不正确")
	}
}

//登录成功后，启动协程 一直监听新消息
func processMes(conn net.Conn) {
	//tf := &utils.Transfer{
	//	Conn: conn,
	//}
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端client等待读取服务器server发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg(),err=", err)
			return
		}
		//读取到了消息
		fmt.Printf("读到了：%v\n", mes)
		//var msg common.Message
		//switch mes.Type {
		//case LoginMesType:
		//
		//}
	}

}
