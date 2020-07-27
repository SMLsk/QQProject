package processor

import (
	"QQ/application/client/model"
	"os"

	// "QQ/application/common"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

type MainProcessor struct {
	Window *window.Window
	Root   *sciter.Element
	Conn   net.Conn
}

var (
	MyMainProcessor *MainProcessor
)

func (this *MainProcessor) DefFunc() {
	fmt.Println()
	this.setCallBackHandler()
	this.setElementHandler()
	this.setWinHandler()
}

func (this *MainProcessor) setElementHandler() {
	closeButton, _ := this.Root.SelectById("CloseButton")
	closeButton.DefineMethod("close", func(args ...*sciter.Value) *sciter.Value {
		os.Exit(1)
		return sciter.NewValue("0")
	})
}

func (this *MainProcessor) setCallBackHandler() {

}

func (this *MainProcessor) setWinHandler() {
	this.Window.DefineFunction("GetFriendsData", func(args ...*sciter.Value) *sciter.Value {
		data, _ := json.Marshal(model.MyFriendsManager)
		fmt.Println(string(data))
		return sciter.NewValue(string(data))
	})
	this.Window.DefineFunction("GetUserInfo", func(args ...*sciter.Value) *sciter.Value {
		user, _ := json.Marshal(model.MyUserModel.UserInfo)
		fmt.Println(string(user))
		return sciter.NewValue(string(user))
	})
	this.Window.DefineFunction("CreateDialog", func(args ...*sciter.Value) *sciter.Value {
		dialogType, _ := strconv.Atoi(args[0].String())
		id, _ := strconv.Atoi(args[1].String())
		MyDialogProcessor.NewDialogBox(dialogType, id)
		return sciter.NewValue("0")
	})
}
