package main

import (
	"QQ/application/server/conf"
	"QQ/application/server/dao"
	"QQ/application/server/model"
	"QQ/application/server/processor"
	"net"
)
import (
	"fmt"
)

func init() {
	dao.MyDao = dao.NewDao()
	processor.MyUserManager = processor.NewUserManager()
	processor.MySmsProcessor = &processor.SmsProcessor{}
	model.MyTemporaryMessage = &model.TemporaryMessage{}
}

func main() {
	listen, err := net.Listen("tcp", conf.Address)
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(conn.RemoteAddr().String(), "连接成功")
		mainProcess := &MainProcess{
			Conn: conn,
		}
		go mainProcess.Run()
	}
}
