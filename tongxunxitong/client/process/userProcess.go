package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "go_code/project01/main/tongxunxitong/common/massage"
	"go_code/project01/main/tongxunxitong/server/utils"
	"net"
	"os"
)

type UserProcess struct {
}

// 处理注册
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	defer conn.Close()
	//2.通过conn发送消息到服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个registerMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
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
	//创建一个Tsansfer实例
	tf := &utils.Tsansfer{
		Conn: conn,
	}
	//发送data到服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err=", err)
	}

	mes, err = tf.ReadPkg() //mes就是RegisterResMes
	if err != nil {
		fmt.Println("tf.ReadPkg err=", err)
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

// 处理登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

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
	tf := &utils.Tsansfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
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

		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UseStatus = message.UserOnline
		//显示当前在线用户ID
		fmt.Println("当前在线用户列表如下")
		for _, v := range loginResMes.UserIds {
			fmt.Println("用户ID:\t", v)
			//完成对客户端onlineUsers的初始化
			user := &message.User{
				UserId:    v,
				UseStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		//创建一个协程，保持服务端和客户端的持续通讯
		go processServerMes(conn)
		//展示次级列表
		for {
			showMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
