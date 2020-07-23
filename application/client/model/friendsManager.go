package model

import (
	"QQ/application/common/message"
	"strconv"
)

type FriendsManager struct {
	Groups map[int]string `json:"groups"`
	//GroupsFriends的切片类型本来应该是int，但是序列化之后，这里还是int，而Friends里面的int已经变为了string，为了方便查询这里也改成string
	//不确定会产生什么风险，不过后面的得注意一些类型转换的问题了
	GroupsFriends map[int][]string           `json:"groupsFriends"`
	Friends       map[int]message.FriendInfo `json:"friends"`
}

var (
	MyFriendsManager *FriendsManager
)

func NewFriendsManager(groups []message.GroupInfo, friends []message.FriendInfo) (myFriendsManager *FriendsManager) {
	myFriendsManager = &FriendsManager{
		Groups:        make(map[int]string, 0),
		GroupsFriends: make(map[int][]string, 0),
		Friends:       make(map[int]message.FriendInfo, 0),
	}
	photoModel := &PhotoModel{}
	for _, val := range groups {
		myFriendsManager.Groups[val.GroupId] = val.GroupName
		myFriendsManager.GroupsFriends[val.GroupId] = make([]string, 0)
	}
	for _, val := range friends {
		val.PhotoAddress = photoModel.DownPhoto(val.PhotoAddress)
		myFriendsManager.GroupsFriends[val.GroupId] = append(myFriendsManager.GroupsFriends[val.GroupId], strconv.Itoa(val.FriendQQ))
		myFriendsManager.Friends[val.FriendQQ] = val
	}

	return
}

func (this *FriendsManager) UpdateStatus(qq int, status int) {
	temp := this.Friends[qq]
	temp.Status = status
	delete(this.Friends, qq)
	this.Friends[qq] = temp
}
