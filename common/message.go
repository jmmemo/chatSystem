package common

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

//消息有类型，和数据本身
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//发送的消息有id和密码
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

//返回的数据，状态码和err
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type RegisterMes struct {
	//
}
