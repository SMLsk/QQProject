package processor

import (
	"QQ/application/common/message"
	"QQ/application/common/utils"
	"QQ/application/server/model"
	"encoding/json"
	"strings"

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
	data, err := json.Marshal(mes)
	if shortMessageMes.Type == 0 {
		recipient, err := MyUserManager.GetOnlineUserByQQ(shortMessageMes.RecipientId)
		if err != nil {
			mesString := string(data)
			mesString = strings.ReplaceAll(mesString, `\`, `\\`)
			model.MyTemporaryMessage.InsertMessage(shortMessageMes.RecipientId, mesString)
			fmt.Println(shortMessageMes.RecipientId, "未上线，转入数据库")
		} else {
			transfer := &utils.Transfer{
				Conn: recipient.Conn,
			}
			err = transfer.WritePkg(data)
			return err
		}
	}
	return
}
