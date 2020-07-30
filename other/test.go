package main

import (
	"QQ/application/common/message"
	"QQ/application/server/dao"
	"QQ/application/server/model"
	"encoding/json"
	"fmt"
	"strings"
	// "strings"
)

func main() {
	startWindow()
}

func startWindow() {
	// fmt.Println("")
	// Window, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{300, 100, 950, 630})
	// if err != nil {
	// 	// log.Fatal(err)
	// }
	// Window.LoadFile("E://Technology//Project//MyGo//src//QQ//application//client//view//dialog0Test.html")
	// Window.SetTitle("Main")
	// Window.Show()
	// Window.Run()
	fmt.Println("")
	dao.MyDao = dao.NewDao()
	MyTemporaryMessage := &model.TemporaryMessage{}
	// fmt.Println()
	messages, _ := MyTemporaryMessage.GetMessageByqq(1234567892)
	for i := 0; i < len(messages); i++ {
		temp := messages[i]
		fmt.Println(temp)
		var mes message.Message
		err := json.Unmarshal(temp, &mes)
		if err != nil {
			fmt.Println(err)
		}
	}

	a := `{"type":"ShortMessageMes","data":"{\"type\":0,\"senderId\":1234567890,\"recipient\":1234567894,\"content\":\"宋魁\",\"rows\":2,\"length\":4,\"chartGroupId\":0,\"time\":1596095167}"}`
	a = strings.ReplaceAll(a, `\`, `\\`)
	fmt.Println(a)
	// var mes message.Message
	// err := json.Unmarshal([]byte(temp), &mes)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
