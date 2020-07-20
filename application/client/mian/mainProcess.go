package main

import (
	"QQ/application/client/model"
	"QQ/application/client/processor"
	"log"
	"net"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

type MainProcess struct {
	Window *window.Window
	Conn   net.Conn
}

func (this *MainProcess) Run() {
	var err error
	this.Window, err = window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{590, 300, 950, 570})
	if err != nil {
		log.Fatal(err)
	}
	this.Window.LoadFile("E://Technology//Project//MyGo//src//QQ//application//client//view//login.html")
	this.Window.SetTitle("登录")
	root, err := this.Window.GetRootElement()

	this.Conn, err = net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		this.Window.Call("error", sciter.NewValue("网络问题，连接服务器失败"))
	}

	loginProcessor := &processor.LoginProcessor{
		Window: this.Window,
		Root:   root,
		Conn:   this.Conn,
	}
	loginProcessor.DefFunc()

	this.Window.Show()
	this.Window.Run()
	if model.MyFriendsManager == nil {
		return
	}
	Window, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{1190, 70, 1480, 630})
	if err != nil {
		// log.Fatal(err)
	}
	Window.LoadFile("E://Technology//Project//MyGo//src//QQ//application//client//view//main.html")
	Window.SetTitle("Main")
	Window.Show()
	Window.Run()
}
