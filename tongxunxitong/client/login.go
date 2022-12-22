package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/project01/main/tongxunxitong/common/massage"
	"net"
)

func login(userId int, userPwd string) (err error) {

	/*fmt.Printf("userId = %d , userPwd = %s", userId, userPwd)
	return nil*/

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	defer conn.Close()

	//2.通过conn发送消息到服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个loginMes结构体
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(LoginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.将data赋给mes.Dta
	mes.Data = string(data)

	//6.将mes.Dta序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes.Dta err=", err)
		return
	}

	//7.这个data就是要发送的消息
	//先发送消息的长度，要把长度信息转为切片
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
	//fmt.Println("客户端发送消息长度成功")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	//接受服务器端返回的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}
	//将mes反序列化为loginResMes并输出提示信息
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
