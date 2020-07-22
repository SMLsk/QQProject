package model

import (
	"QQ/application/common"
	"QQ/application/common/message"
	"QQ/application/common/utils"
	"encoding/json"

	"errors"
	"fmt"
	"net"
)

type UserModel struct {
	UserInfo message.User `json:"user"`
}

var (
	MyUserModel *UserModel
)

func Login(qq int, pwd string, conn net.Conn) (myUserModel *UserModel, err error) {
	var loginMes message.LoginMes
	loginMes.User.QQ = qq
	loginMes.User.Pwd = common.Md5(pwd)

	data, err := common.AssembleMes(message.LoginMesType, loginMes)
	transfer := &utils.Transfer{
		Conn: conn,
	}
	err = transfer.WritePkg(data)

	resMes, err := transfer.ReadPkg()
	if err != nil {
		return
	}
	// fmt.Println("接收成功")
	var loginRes message.LoginRes
	err = json.Unmarshal([]byte(resMes.Data), &loginRes)
	if loginRes.Code == 200 {
		//myUserModel.UserInfo = loginRes.UserInfo
		photoModel := &PhotoModel{}
		myUserModel = &UserModel{
			UserInfo: loginRes.UserInfo,
		}
		fmt.Println()
		myUserModel.UserInfo.PhotoAddress = photoModel.DownPhoto(myUserModel.UserInfo.PhotoAddress)
		MyFriendsManager = NewFriendsManager(loginRes.Groups, loginRes.Friends)
		// fmt.Println(MyFriendsManager)
	} else if loginRes.Code == 500 {
		fmt.Println(loginRes.Error)
		return myUserModel, errors.New("Error")
	}
	return
}
