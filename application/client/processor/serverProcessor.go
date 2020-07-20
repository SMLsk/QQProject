package processor

import (
	"QQ/application/common/message"
	"QQ/application/common/utils"

	// "encoding/json"
	"fmt"
	"net"
)

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例, 不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		fmt.Println(1, mes)
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑
		// fmt.Printf("mes=%v\n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			fmt.Println(mes)
		default:
			fmt.Println("ERROR")
		}

	}
}
