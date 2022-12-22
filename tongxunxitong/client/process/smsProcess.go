package process

import (
	"encoding/json"
	"fmt"
	_ "go_code/project01/main/tongxunxitong/client/modle"
	message "go_code/project01/main/tongxunxitong/common/massage"
	"go_code/project01/main/tongxunxitong/server/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMse(content string) (err error) {
	//1.创建一个mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = curUser.UserId
	smsMes.UseStatus = curUser.UseStatus

	//3.将要发送的消息序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送消息
	tf := &utils.Tsansfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err=", err)
		return
	}
	return
}
