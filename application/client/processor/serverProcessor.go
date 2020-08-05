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
		fmt.Println("客户端正在等待读取服务器发送的消息")
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
	fmt.Println("进入处理")
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
		var shortMessageMes message.ShortMessageMes
		json.Unmarshal([]byte(mes.Data), &shortMessageMes)
		//判断该消息的来源好友的对话框是否存在

		//这里应该判断一下是群消息还是个人消息，如果是群消息，就通过ChatGroupId来检索，否则通过SenderId检索
		var chartId int
		if shortMessageMes.Type == 0 {
			chartId = shortMessageMes.SenderId
		} else {
			chartId = shortMessageMes.ChartGroupId
		}
		//若存在，则直接调用对应对话框，并推送
		var ok bool
		var dialog *Dialog
		dialog, ok = MyDialogProcessor.Dialogs[chartId]
		if ok {
			dialog.Window.Call("add", sciter.NewValue(mes.Data))
		}
		model.MySmsListManager.SaveSms(chartId, shortMessageMes)
		fmt.Println(model.MySmsListManager.SmsList)
		//若不存在，则存入消息列表中，在存入函数中，调用界面添加消息的按钮

	default:
		fmt.Println("ERROR")
	}
}
