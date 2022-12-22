package modle

import (
	message "go_code/project01/main/tongxunxitong/common/massage"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
