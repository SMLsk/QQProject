package processor

import (
	"QQ/application/common/message"
	"QQ/application/common/utils"
	"encoding/json"

	// "errors"
	"fmt"
)

type SmsProcessor struct {
}

var (
	MySmsProcessor *SmsProcessor
)

func (this *SmsProcessor) Transpond(mes message.Message) (err error) {
	fmt.Println("")
	var shortMessageMes message.ShortMessageMes
	err = json.Unmarshal([]byte(mes.Data), &shortMessageMes)
	if shortMessageMes.Type == 0 {
		recipient, err := MyUserManager.GetOnlineUserByQQ(shortMessageMes.RecipientId)
		if err != nil {
			//存入数据库
		} else {
			transfer := &utils.Transfer{
				Conn: recipient.Conn,
			}
			data, err := json.Marshal(mes)
			err = transfer.WritePkg(data)
			return err
		}
	}
	return
}
