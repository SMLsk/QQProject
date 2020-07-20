package model

import (
	"QQ/application/common"
	"QQ/application/common/message"
	"QQ/application/server/dao"
	"fmt"
	"strconv"
)

type FriendsModel struct{}

func (this *FriendsModel) GetFriendsByqq(qq int) (friends []message.FriendInfo, err error) {
	fmt.Println("")
	sql := "select group_id,friend_qq from `qq_friends` where user_qq = " + strconv.Itoa(qq)
	res, err := dao.MyDao.ExecuteDql(sql)
	if err != nil {
		return
	}
	userModel := UserModel{}
	for _, val := range res {
		var friendInfo message.FriendInfo
		friendInfo.FriendQQ = common.Atoi(val["friend_qq"])
		friendInfo.GroupId = common.Atoi(val["group_id"])
		user := userModel.GetUserByqq(friendInfo.FriendQQ)
		friendInfo.NickName = user.NickName
		friendInfo.PhotoID = user.PhotoID
		friendInfo.Sign = user.Sign
		friends = append(friends, friendInfo)
	}
	return
}
