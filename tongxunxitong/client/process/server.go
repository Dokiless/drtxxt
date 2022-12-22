package process

import (
	"encoding/json"
	"fmt"
	message "go_code/project01/main/tongxunxitong/common/massage"
	"go_code/project01/main/tongxunxitong/server/utils"
	"net"
	"os"
)

// 显示登录界面
func showMenu() {
	fmt.Println("-----用户登陆成功-----")
	fmt.Println("1.查看用户列表")
	fmt.Println("2.发送消息")
	fmt.Println("3.查看消息列表")
	fmt.Println("4.退出系统")
	fmt.Println("请选择1-4")
	var key int
	var content string
	sp := &SmsProcess{}
	fmt.Scanf("%d\n", &key)

	switch key {
	case 1:
		//fmt.Println("用户列表")
		outputOnlineUsers()
	case 2:
		fmt.Println("输入您的消息")
		fmt.Scanf("%s\n", &content)
		sp.SendGroupMse(content)
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("您选择退出系统")
		os.Exit(0)
	default:
		fmt.Println("您输入的选项错误，请重新输入")
	}
}

// 保持服务器和客户端的通讯
func processServerMes(conn net.Conn) {
	tf := &utils.Tsansfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待服务器传送消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//fmt.Printf("mes=%v", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: //接收到用户在线状态变更的消息
			var notifyUserStatusMes *message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatusMes(notifyUserStatusMes)
		case message.SmsMesType:
			outputMes(mes)
		default:
			fmt.Println("无法识别的消息")
		}
	}
}
