package utils

import (
	"chatroom/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (msg common.Message, err error) {

	n, err := this.Conn.Read(this.Buf[:4])
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println("this.Conn.Read(this.Buf[:4]),err=", err)
	}
	fmt.Println("从client读取的buf长度是", this.Buf[:n])

	///////////现在已经读到客户端传送的byte数据,要将其反序列化成msg
	var pkgLen uint32 ////////////////////这里开始很重要(服务器server读取客户端client发送的信息的步骤)/////////////////////////////////
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	/////为了读取能够成功，需要拿到 转换成uint32类型 的长度
	_, err = this.Conn.Read(this.Buf[:pkgLen]) ///////然后再按照这个长度，重新读取一遍,还是存到buf里面
	if err != nil {
		fmt.Println("this.Conn.Read(this.Buf[:pkgLen]),err=", err)
		return
	}
	//
	err = json.Unmarshal(this.Buf[:pkgLen], &msg) /////////从buf指定长度，反序列化输出一个msg结构体,这才是能看到的信息
	if err != nil {
		fmt.Println("json.Unmarshal(this.Buf[:pkgLen], &msg),err=", err)
		return
	}
	return
	/////////////
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//上面处理长度
	_, err = this.Conn.Write(this.Buf[:4])
	if err != nil {
		fmt.Println("Conn.Write(this.Buf[:]),err=", err)
		return
	}
	//下面发送数据本身

	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("this.Conn.Write(data),err=", err)
		return
	}
	return
}
