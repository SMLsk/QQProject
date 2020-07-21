package processor

import (
	"QQ/application/client/model"
	"encoding/json"
	"fmt"
	"net"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

type MainProcessor struct {
	Window *window.Window
	Root   *sciter.Element
	Conn   net.Conn
}

func (this *MainProcessor) DefFunc() {
	fmt.Println()
	this.setCallBackHandler()
	this.setElementHandler()
	this.setWinHandler()
}

func (this *MainProcessor) setElementHandler() {

}

func (this *MainProcessor) setCallBackHandler() {

}

func (this *MainProcessor) setWinHandler() {
	this.Window.DefineFunction("GetFriendsData", func(args ...*sciter.Value) *sciter.Value {
		data, _ := json.Marshal(model.MyFriendsManager)
		fmt.Println(string(data))
		return sciter.NewValue(string(data))
	})
}
