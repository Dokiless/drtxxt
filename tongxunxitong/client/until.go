package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	message "go_code/project01/main/tongxunxitong/common/massage"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	//fmt.Println("开始读取客户端发送的消息长度")
	_, err = conn.Read(buf[:4])
	if err != nil {

		err = errors.New("read pkg header error")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	//根据pkgLen读取消息
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	//把pkgLen反序列化成mes

	err = json.Unmarshal(buf[:pkgLen], &mes)

	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 编写writePkg函数
func writePkg(conn net.Conn, data []byte) (err error) {
	//1.先发送消息长度
	var pckLen uint32
	pckLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pckLen)
	//传入消息长度
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	//2.发送消息
	n, err = conn.Write(data)
	if n != int(pckLen) || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	return
}
