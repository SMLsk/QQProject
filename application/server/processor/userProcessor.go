package processor

import (
	"QQ/application/common"
	"QQ/application/common/message"
	"QQ/application/common/utils"
	"QQ/application/server/model"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct {
	Conn    net.Conn
	QQ      int
	Friends []int
}

func (this *UserProcessor) LoginProcess(mes message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("loginRes反序列化失败", err)
		return
	}
	userModel := &model.UserModel{}
	user, err := userModel.Login(loginMes.User.QQ, loginMes.User.Pwd)
	// fmt.Println(2, user)
	var loginRes message.LoginRes

	if err != nil {
		loginRes.Code = 500
		if err == model.ERROR_USER_NOTEXIST {
			loginRes.Error = "用户不存在"
		} else if err == model.ERROR_USER_PWD {
			loginRes.Error = "密码不存在"
		} else {
			loginRes.Error = "未知错误"
		}
	} else {
		this.QQ = user.QQ
		MyUserManager.AddOnlineUser(this)

		fmt.Println(err)

		loginRes = this.assembleLoginSuccessRes(user)
		// fmt.Println(loginRes)
		for _, val := range loginRes.Friends {
			this.Friends = append(this.Friends, val.FriendQQ)
		}
		//这一句必须放在下面，否则尚未获取到friends，将不会发送
		err = this.NotifyOtherOnlineUsers(message.Online)
		if err != nil {
			return
		}
	}

	// fmt.Println(loginRes)
	data, err := common.AssembleMes(message.LoginResType, loginRes)

	transfer := &utils.Transfer{
		Conn: this.Conn,
	}

	err = transfer.WritePkg(data)
	return
}

func (this *UserProcessor) assembleLoginSuccessRes(user message.User) (loginRes message.LoginRes) {
	groupsModel := &model.GroupsModel{}
	friendsModel := &model.FriendsModel{}
	loginRes.Code = 200
	loginRes.UserInfo = user
	loginRes.Groups, _ = groupsModel.GetGroupsByqq(user.QQ)
	loginRes.Friends, _ = friendsModel.GetFriendsByqq(user.QQ)

	for i := 0; i < len(loginRes.Friends); i++ {

		_, err := MyUserManager.GetOnlineUserByQQ(loginRes.Friends[i].FriendQQ)
		if err != nil {
			loginRes.Friends[i].Status = message.Offline
		} else {
			loginRes.Friends[i].Status = message.Online
		}
	}
	fmt.Println(loginRes)
	return loginRes
}

func (this *UserProcessor) NotifyOtherOnlineUsers(status int) (err error) {
	for _, qq := range this.Friends {
		up, err := MyUserManager.GetOnlineUserByQQ(qq)
		if err == nil {
			err = up.NotifyMeUserStatus(this.QQ, status)
			if err != nil {
				return err
			}
		}
	}
	return
}

func (this *UserProcessor) NotifyMeUserStatus(qq int, status int) (err error) {
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.QQ = qq
	notifyUserStatusMes.Status = status

	data, err := common.AssembleMes(message.NotifyUserStatusMesType, notifyUserStatusMes)
	if err != nil {
		return err
	}

	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	fmt.Println("1", err)
	return
}
