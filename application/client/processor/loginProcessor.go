package processor

import (
	"QQ/application/client/model"
	"fmt"
	"strconv"

	// "log"
	"net"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

type LoginProcessor struct {
	Window *window.Window
	Root   *sciter.Element
	Conn   net.Conn
}

func (this *LoginProcessor) DefFunc() {
	fmt.Println()
	this.setCallBackHandler()
	this.setElementHandler()
	this.setWinHandler()
}

func (this *LoginProcessor) setElementHandler() {
	loginButton, _ := this.Root.SelectById("loginButton")
	loginButton.DefineMethod("login", func(args ...*sciter.Value) *sciter.Value {
		qq, _ := strconv.Atoi(args[0].String())
		pwd := args[1].String()
		var err error
		model.MyUserModel, err = model.Login(qq, pwd, this.Conn)
		if err != nil {
			fmt.Println(err)
			return sciter.NewValue("error")
		}
		// fmt.Println(model.MyFriendsManager)
		serverProcessor := &ServerProcessor{}
		go serverProcessor.receiveMes(this.Conn)
		return sciter.NewValue("ok")
	})

}

func (this *LoginProcessor) setCallBackHandler() {

}

func (this *LoginProcessor) setWinHandler() {

}
