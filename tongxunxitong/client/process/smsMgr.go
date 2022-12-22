package process

import (
	"encoding/json"
	"fmt"
	message "go_code/project01/main/tongxunxitong/common/massage"
)

func outputMes(mes message.Message) {
	var smsMes message.SmsMes
	json.Unmarshal([]byte(mes.Data), &smsMes)
	message := fmt.Sprintf("用户:\t%d 对大家说\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(message)
	fmt.Print("\n")
}
