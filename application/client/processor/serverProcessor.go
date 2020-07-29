package processor

import (
	"QQ/application/client/model"
	"encoding/json"

	// "QQ/application/client/model"
	"QQ/application/common/message"
	"QQ/application/common/utils"

	"github.com/sciter-sdk/go-sciter"

	// "encoding/json"
	"fmt"
	"net"
)

type ServerProcessor struct{}

//和服务器保持通讯
func (this *ServerProcessor) receiveMes(conn net.Conn) {
	//创建一个transfer实例, 不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		// fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑
		this.serverProcessMes(mes)
	}
}

func (this *ServerProcessor) serverProcessMes(mes message.Message) {
	// fmt.Printf("mes=%v\n", mes)
	switch mes.Type {
	case message.NotifyUserStatusMesType:
		var notifyUserStatusMes message.NotifyUserStatusMes
		err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
		if err != nil {
			fmt.Println(err)
		}
		model.MyFriendsManager.UpdateStatus(notifyUserStatusMes.QQ, notifyUserStatusMes.Status)
		data, _ := json.Marshal(model.MyFriendsManager)
		MyMainProcessor.Window.Call("initGroups", sciter.NewValue(string(data)))
		fmt.Println(mes)
	case message.ShortMessageMesType:
		fmt.Println(mes)
	default:
		fmt.Println("ERROR")
	}
}
