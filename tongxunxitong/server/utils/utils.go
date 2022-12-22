package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	message "go_code/project01/main/tongxunxitong/common/massage"
	"net"
)

// 这里将方法封装到结构体中
type Tsansfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Tsansfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 8096)
	//fmt.Println("开始读取客户端发送的消息长度")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		err = errors.New("read pkg header error")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkgLen读取消息
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	//把pkgLen反序列化成mes

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)

	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 编写WritePkg函数
func (this *Tsansfer) WritePkg(data []byte) (err error) {
	//1.先发送消息长度
	var pckLen uint32
	pckLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pckLen)
	//传入消息长度
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	//2.发送消息
	n, err = this.Conn.Write(data)
	if n != int(pckLen) || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	return
}
