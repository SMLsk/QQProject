package model

import (
	"QQ/application/common/message"
)

type FriendsManager struct {
	Groups        map[int]string               `json:"groups"`
	GroupsFriends map[int][]message.FriendInfo `json:"groupsFriends"`
}

var (
	MyFriendsManager *FriendsManager
)

func NewFriendsManager(groups []message.GroupInfo, friends []message.FriendInfo) (myFriendsManager *FriendsManager) {
	myFriendsManager = &FriendsManager{
		Groups:        make(map[int]string, 0),
		GroupsFriends: make(map[int][]message.FriendInfo, 0),
	}

	for _, val := range groups {
		myFriendsManager.Groups[val.GroupId] = val.GroupName
		myFriendsManager.GroupsFriends[val.GroupId] = make([]message.FriendInfo, 0)
	}
	for _, val := range friends {
		myFriendsManager.GroupsFriends[val.GroupId] = append(myFriendsManager.GroupsFriends[val.GroupId], val)
	}

	return
}
