package processor

import (
	"QQ/application/common/message"
	"errors"
	"fmt"
)

type UserManager struct {
	onlineUsers map[int]*UserProcessor
}

var (
	MyUserManager *UserManager
)

func NewUserManager() (myUserManager *UserManager) {
	myUserManager = &UserManager{
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
	return
}

func (this *UserManager) AddOnlineUser(userProcessor *UserProcessor) {
	this.onlineUsers[userProcessor.QQ] = userProcessor
}

func (this *UserManager) DelOnlineUser(qq int) {

	err := this.onlineUsers[qq].NotifyOtherOnlineUsers(message.Offline)
	if err != nil {
		fmt.Println(err)
	}
	delete(this.onlineUsers, qq)
}

func (this *UserManager) GetAllOnlineUser() map[int]*UserProcessor {
	return this.onlineUsers
}

func (this *UserManager) GetOnlineUserByQQ(qq int) (up *UserProcessor, err error) {
	up, ok := this.onlineUsers[qq]
	if !ok {
		err = errors.New("用户不存在")
		return
	}
	return
}
