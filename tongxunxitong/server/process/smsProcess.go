package process2

import (
	"encoding/json"
	"fmt"
	message "go_code/project01/main/tongxunxitong/common/massage"
	"go_code/project01/main/tongxunxitong/server/utils"
	"net"
)

type SmsProcess struct {
}

// 编写方法转发消息
func (this *SmsProcess) SendGroupMes(mes message.Message) {
	//取出mes的内容
	var smsMes message.Message
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	for _, up := range userMgr.onlineUsers {
		this.SendMes(data, up.Conn)
	}
}

func (this *SmsProcess) SendMes(data []byte, conn net.Conn) {
	//创建一个Tansfer实例,用于发送消息
	tf := &utils.Tsansfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err=", err)
		return
	}
}
