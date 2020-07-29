package main

import (
	"QQ/application/common/message"
	"QQ/application/common/utils"
	"QQ/application/server/processor"
	"fmt"
	"io"
	"net"
)

type MainProcess struct {
	Conn net.Conn
	QQ   int
}

func (this *MainProcess) Run() {
	defer this.Conn.Close()
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	for {
		mes, err := transfer.ReadPkg()
		if err != nil {

			if err == io.EOF {

				break
			}
			continue
		}
		err = this.MesProcess(mes)
		if err != nil {
			fmt.Println(err)
		}
	}
	processor.MyUserManager.DelOnlineUser(this.QQ)
	fmt.Println(this.Conn.RemoteAddr().String(), "连接断开")
}

func (this *MainProcess) MesProcess(mes message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		fmt.Println("成功获取type")
		userProcess := &processor.UserProcessor{
			Conn: this.Conn,
		}
		err = userProcess.LoginProcess(mes)
		if err != nil {
			return
		}
		this.QQ = userProcess.QQ
		return
	case message.RegisterMesType:
		// userProcess := &process.UserProcess{
		// 	Conn: this.Conn,
		// }
		// err = userProcess.ServerProcessRegister(mes)
		// if err != nil {
		// 	return
		// }
	case message.ShortMessageMesType:
		fmt.Println(mes)
		err = processor.MySmsProcessor.Transpond(mes)
		return
	default:
		fmt.Println("未知类型信息")
	}
	return
}
