package model

import (
	"QQ/application/common"
	// "QQ/application/common"
	"QQ/application/common/message"
	"QQ/application/server/dao"
	"fmt"
	"strconv"
)

type UserModel struct {
}

func (this *UserModel) Login(qq int, pwd string) (user message.User, err error) {
	sql := "select pwd from qq_user where qq = " + strconv.Itoa(qq)
	fmt.Println("")
	res, err := dao.MyDao.ExecuteDqlOne(sql)
	// fmt.Println(err)
	if err != nil {
		err = ERROR_USER_NOTEXIST
		return
	}
	if pwd == res["pwd"] {
		user = this.GetUserByqq(qq)
		return
	} else {
		err = ERROR_USER_PWD
	}
	return
}

func (this *UserModel) GetUserByqq(qq int) (user message.User) {
	sql := "select * from qq_user where qq = " + strconv.Itoa(qq)

	res, err := dao.MyDao.ExecuteDqlOne(sql)
	if err != nil {
		return
	}
	user = message.User{
		QQ:        qq,
		Sign:      res["sign"],
		PhotoID:   common.Atoi(res["photoID"]),
		NickName:  res["nickName"],
		Sex:       res["sex"],
		Birthday:  res["birthday"],
		BloodType: res["bloodType"],
		Diploma:   res["diploma"],
		Telephone: res["telephone"],
		Email:     res["email"],
		Address:   res["address"],
	}
	return
}
