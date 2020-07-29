package model

import (
	"QQ/application/common"
	"QQ/application/common/message"
	"QQ/application/common/utils"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"unicode"
	// "strconv"
)

type SmsModel struct {
	Conn     net.Conn
	QQ       int
	Transfer *utils.Transfer
}

var (
	MySmsModel *SmsModel
)

func NewSmsModel(conn net.Conn, qq int) (mySmsModel *SmsModel) {
	mySmsModel = &SmsModel{
		Conn:     conn,
		QQ:       qq,
		Transfer: &utils.Transfer{Conn: conn},
	}
	return
}

func (this *SmsModel) SendSms(dialogType int, recipientId int, content string) (messageString string, err error) {
	fmt.Println(dialogType, recipientId, content)
	var shortMessageMes message.ShortMessageMes
	shortMessageMes.Type = dialogType
	shortMessageMes.Content = content
	shortMessageMes.SenderId = this.QQ
	shortMessageMes.RecipientId = recipientId
	shortMessageMes.Rows, shortMessageMes.Length = this.countRows(content)
	data, err := common.AssembleMes(message.ShortMessageMesType, shortMessageMes)
	err = this.Transfer.WritePkg(data)
	data, err = json.Marshal(shortMessageMes)
	messageString = string(data)
	return
}

func (this *SmsModel) countRows(content string) (rows float64, length float64) {
	for _, c := range content {
		if unicode.Is(unicode.Han, c) {
			length += 2
		} else {
			length += 1
		}
	}
	rows = math.Ceil(length/23 + 1)
	return
}

func (this *SmsModel) ReciveMessage(mes message.Message) {
	//接收消息，并反序列化

	//
}
