package common

import (
	"QQ/application/common/message"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strconv"
)

func AssembleMes(mesType string, dataStruct interface{}) (data []byte, err error) {
	var mes message.Message
	mes.Type = mesType

	data, err = json.Marshal(dataStruct)
	if err != nil {
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	return
}

func Md5(str string) (md5Str string) {
	res := md5.Sum([]byte(str))
	md5Str = fmt.Sprintf("%x", res)

	return
}

func Atoi(str string) (i int) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	return
}