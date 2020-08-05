package model

import (
	"QQ/application/common/message"
	_ "fmt"
)

type SmsListManager struct {
	//key存储消息来源的id
	SmsList map[int][]message.ShortMessageMes
}

var (
	MySmsListManager *SmsListManager
)

func NewSmsListManager() (mySmsListManager *SmsListManager) {
	mySmsListManager = &SmsListManager{
		SmsList: make(map[int][]message.ShortMessageMes, 0),
	}
	return
}

func (this *SmsListManager) SaveSms(id int, shortMessageMes message.ShortMessageMes) {

	this.SmsList[id] = append(this.SmsList[id], shortMessageMes)

}
