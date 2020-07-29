package model

import (
	"QQ/application/common"
	"QQ/application/common/message"
	"QQ/application/common/utils"
	"fmt"
	"net"
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

func (this *SmsModel) SendSms(dialogType int, recipientId int, content string) (err error) {
	fmt.Println(dialogType, recipientId, content)
	var shortMessageMes message.ShortMessageMes
	shortMessageMes.Type = dialogType
	shortMessageMes.Content = content
	shortMessageMes.SenderId = this.QQ
	shortMessageMes.RecipientId = recipientId
	data, err := common.AssembleMes(message.ShortMessageMesType, shortMessageMes)

	err = this.Transfer.WritePkg(data)
	return err
}
