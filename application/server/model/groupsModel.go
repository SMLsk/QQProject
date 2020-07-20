package model

import (
	"QQ/application/common"
	"QQ/application/common/message"
	"QQ/application/server/dao"
	"fmt"
	"strconv"
)

type GroupsModel struct{}

func (this *GroupsModel) GetGroupsByqq(qq int) (groups []message.GroupInfo, err error) {
	fmt.Println("")
	sql := "select id,group_name from `qq_groups` where user_qq = " + strconv.Itoa(qq)
	res, err := dao.MyDao.ExecuteDql(sql)
	if err != nil {
		return
	}
	for _, val := range res {
		var groupInfo message.GroupInfo

		groupInfo.GroupId = common.Atoi(val["id"])
		groupInfo.GroupName = val["group_name"]
		groups = append(groups, groupInfo)
	}
	return
}
