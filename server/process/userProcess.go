package process

import (
	"chatroom/common"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//登录
func (this *UserProcess) ServerProcessLogin(msg *common.Message) (err error) {
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(msg.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(msg.Data), loginMes),err=", err)
		return
	}
	//构建返回消息
	var resMes common.Message
	resMes.Type = common.LoginResMesType

	var loginResMes common.LoginResMes
	///////////////////

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	///////////////////
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "未知错误，服务器内部错误"
		}
		//loginResMes.Code = 500
		//loginResMes.Error = "用户不存在，请注册再使用"
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功!!!!!!!")
	}

	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	//合法
	//	loginResMes.Code = 200
	//} else {
	//	//不合法
	//	loginResMes.Code = 500
	//	loginResMes.Error = "用户不存在，请注册再使用"
	//}
	/////////////////////
	//为返回resMes做准备
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes),err=", err)
		return
	}
	resMes.Data = string(data)
	//序列化准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes),err=", err)
		return
	}
	/////writePkg
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data),err=", err)
		return
	}
	return
}
