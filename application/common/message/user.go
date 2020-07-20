package message

type User struct {
	QQ        int    `json"qq"`
	Pwd       string `json"pwd"`
	Sign      string `json"sign"`
	PhotoID   int    `json:"photoID"`
	NickName  string `json:"nickName"`
	Sex       string `json:"sex"`
	Birthday  string `json:"birthday"`
	BloodType string `json:"bloodType"`
	Diploma   string `json:"diploma"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

type FriendInfo struct {
	FriendQQ int    `json"qq"`
	Sign     string `json"sign"`
	PhotoID  int    `json:"photoID"`
	NickName string `json:"nickName"`
	GroupId  int    `json:"groupId"`
	Status   int    `json:"status"`
}

type GroupInfo struct {
	GroupId   int    `json"groupId"`
	GroupName string `json"groupName"`
}
