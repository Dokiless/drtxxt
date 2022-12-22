package main

import (
	"fmt"
	"go_code/project01/main/tongxunxitong/server/model"
	"net"
	"time"
)

/*func readPkg(conn net.Conn) (mes message.Message, err error) {

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
}*/

/*// 编写一个serverProcessloginMes函数,专门处理登录消息
func serverProcessloginMe(conn net.Conn, mes message.Message) (err error) {
	//核心代码
	//从mes中取出mes.Data,再将其反序列化
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginmes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2.再声明一个LoginResMes并完成赋值
	var loginResMes message.LoginResMes
	//先假定用户ID为111，密码为abc,进行判断
	if loginmes.UserId == 111 && loginmes.UserPwd == "abc" {
		loginResMes.Code = 200 //状态码为200定义为表示登陆成功
	} else {
		loginResMes.Code = 500 //状态码为500定义为表示登陆失败
		loginResMes.Error = "您输入的账号或者密码不正确，请重新输入"
	}

	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.将resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//6.发送消息,将代码封装到writePkg函数中
	err = writePkg(conn, data)
	return
}*/

/*// 编写一个serverProcessMes函数，来根据不同的消息类型调用对应的函数做出处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		serverProcessloginMe(conn, *mes)
	case message.RegisterMesType:
		//处理注册
	}
	return
}*/

// 处理和客户端的通讯
func pocess(conn net.Conn) {
	defer conn.Close()
	//调用总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("协程出现问题")
	}
}

// 这里编写一个函数，完成对UserDao的初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(*pool)
}

func main() {

	//当服务器启动，就初始化redis连接池
	initpool("localhost:6379", 16, 0, 100*time.Second)
	initUserDao()

	//提示监听8889端口
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)

	}
	//监听成功等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("listen.Accept err =", err)

		}
		//一旦连接成功，就启用一个协程与客户端保持通讯
		go pocess(conn)
	}
}
