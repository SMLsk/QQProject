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

func Atoi(str interface{}) (i int) {
	if str2, ok := str.(string); ok {
		var err error
		i, err = strconv.Atoi(str2)
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		i = str.(int)
		return
	}

}

func ByteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
