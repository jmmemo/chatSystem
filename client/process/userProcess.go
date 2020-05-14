package process

import (
	"chatroom/client/utils"
	"chatroom/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial(8889),err=", err)
		return
	}

	//
	var msg common.Message
	msg.Type = common.LoginMesType

	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes),err=", err)
		return
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal(msg),err=", err)
		return
	}

	/////////////
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//////////
	_, err = conn.Write(buf[:])
	if err != nil {
		fmt.Println("conn.Write(buf[:]),err=", err)
		return
	}
	fmt.Println("client发送的数据是,长度是", string(data), len(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data),err=", err)
		return
	}

	///////////
	tf := &utils.Transfer{
		Conn: conn,
	}
	//////
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn),err=", err)
		return
	}
	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(msg.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(msg.Data), &loginResMes),err=", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("登录成功(来自client的userProcess)")
		//////////////////
		go processMes(conn) ///////////读他妈的
		///////////////////
		for {
			showMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	///////
	return
}
