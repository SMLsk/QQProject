package model

import (
	"QQ/application/server/dao"
	"fmt"
	"strconv"
)

type PhotoModel struct {
}

func (this *PhotoModel) GetPhotoAddressById(id int) (address string) {
	sql := "select address from `qq_photo` where id = " + strconv.Itoa(id)
	res, err := dao.MyDao.ExecuteDqlOne(sql)
	// fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	return res["address"]
}
