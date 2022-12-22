package main

import (
	"fmt"
	message "go_code/project01/main/tongxunxitong/common/massage"
	process2 "go_code/project01/main/tongxunxitong/server/process"
	"go_code/project01/main/tongxunxitong/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMes函数，来根据不同的消息类型调用对应的函数做出处理
func (this *Processor) serverProcessMes(mes message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		pro := &process2.UserProcess{
			Conn: this.Conn,
		}
		pro.ServerProcessloginMes(mes)
	case message.RegisterMesType:
		//处理注册
		pro := &process2.UserProcess{
			Conn: this.Conn,
		}
		pro.ServerProcessRegisterMes(mes)
	case message.SmsMesType:
		sms := &process2.SmsProcess{}
		sms.SendGroupMes(mes)
	}
	return
}

func (this *Processor) process2() (err error) {
	for {
		//封装一个函数
		tf := &utils.Tsansfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg err=", err)
			return err
		}
		fmt.Println("mes=", mes)

		//调用serverProcessMes函数
		err = this.serverProcessMes(mes)
		if err != nil {
			fmt.Println("serverProcessMes err=", err)
			return err
		}

	}
}
