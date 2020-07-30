package processor

import (
	"QQ/application/client/model"
	"encoding/json"
	"fmt"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

type DialogProcessor struct {
	Dialogs map[int]*Dialog
}

type Dialog struct {
	Type   int
	Id     int
	Window *window.Window
	Root   *sciter.Element
}

var (
	MyDialogProcessor *DialogProcessor
)

func init() {
	MyDialogProcessor = NewDialogProcessor()
}

func NewDialogProcessor() (dialogProcessor *DialogProcessor) {
	dialogProcessor = &DialogProcessor{
		Dialogs: make(map[int]*Dialog, 0),
	}
	return
}

//这里想到以后可能有群聊，可以用同一个函数，通过不同的参数来区分，用不定参数列表或者直接确定呢？
func (this *DialogProcessor) NewDialogBox(dialogType int, id int) {
	//fmt.Printf("%T,%v,%T,%v", dialogType, dialogType, id, id)
	//用dialogType来区分是群聊还是私聊，0代表私聊，1代表群聊
	dialogWindow, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{300, 100, 950, 630})
	if err != nil {
		fmt.Println(err)
	}
	if dialogType == 0 {

		dialogWindow.LoadFile("E://Technology//Project//MyGo//src//QQ//application//client//view//dialog0.html")
	} else {
		fmt.Println("开启群聊")
		dialogWindow.LoadFile("./test.html")
	}
	root, _ := dialogWindow.GetRootElement()
	dialog := &Dialog{
		Type:   dialogType,
		Id:     id,
		Window: dialogWindow,
		Root:   root,
	}
	dialog.DefFunc()
	this.Dialogs[id] = dialog
	dialogWindow.Show()
	dialogWindow.Run()
}

func (this *Dialog) DefFunc() {
	this.setCallBackHandler()
	this.setElementHandler()
	this.setWinHandler()
}

func (this *Dialog) setElementHandler() {
	sendButton, _ := this.Root.SelectById("sendButton")
	sendButton.DefineMethod("sendMessage", func(args ...*sciter.Value) *sciter.Value {
		//fmt.Println("ok")
		content := args[0].String()
		messageString, err := model.MySmsModel.SendSms(this.Type, this.Id, content)
		//调用SmsModel的方法发送消息
		if err != nil {
			fmt.Println(err)
		}
		this.Window.Call("add", sciter.NewValue(messageString))
		return sciter.NewValue(0)
	})
}

func (this *Dialog) setCallBackHandler() {

}

func (this *Dialog) setWinHandler() {
	//判断为好友对话之后，通过id获取好友信息，并传给前端
	this.Window.DefineFunction("GetdialogInfo", func(args ...*sciter.Value) *sciter.Value {
		//model.MyFriendsManager.Friends[this.Id]
		data, _ := json.Marshal(model.MyFriendsManager.Friends[this.Id])
		fmt.Println(string(data))
		return sciter.NewValue(string(data))
	})
	this.Window.DefineFunction("GetUserInfo", func(args ...*sciter.Value) *sciter.Value {
		user, _ := json.Marshal(model.MyUserModel.UserInfo)
		fmt.Println(string(user))
		return sciter.NewValue(string(user))
	})
}
