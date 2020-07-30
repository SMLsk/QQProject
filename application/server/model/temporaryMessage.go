package model

import (
	"QQ/application/server/dao"
	"fmt"
	"strconv"
	// "strings"
)

type TemporaryMessage struct {
}

var (
	MyTemporaryMessage *TemporaryMessage
)

func (this *TemporaryMessage) GetMessageByqq(qq int) (messages [][]byte, err error) {
	fmt.Println("")
	sql := "select message from `qq_temporary_message` where recipient = " + strconv.Itoa(qq)
	res, err := dao.MyDao.ExecuteDql(sql)
	if err != nil {
		return
	}
	for _, val := range res {
		fmt.Println(val["message"])
		messages = append(messages, []byte(val["message"]))
	}
	this.DeleteMessage(qq)
	//fmt.Println(messages)
	return
}

func (this *TemporaryMessage) InsertMessage(qq int, message string) {
	sql := "insert into `qq_temporary_message` value(" + strconv.Itoa(qq) + ",'" + message + "')"
	fmt.Println(sql)
	_, err := dao.MyDao.ExecuteDml(sql)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(res)
}

func (this *TemporaryMessage) DeleteMessage(qq int) {
	sql := "delete from `qq_temporary_message` where recipient = " + strconv.Itoa(qq)
	fmt.Println(sql)
	res, err := dao.MyDao.ExecuteDml(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res = ", res)
}
