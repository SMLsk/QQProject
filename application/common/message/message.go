package message

const (
	LoginMesType            = "LoginMes"
	LoginResType            = "LoginRes"
	RegisterMesType         = "RegisterMes"
	RegisterResType         = "RegisterRes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

const (
	Online = iota
	Offline
	Busy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	User User `json:"user"`
}

type LoginRes struct {
	Code     int          `json:"code"`
	Error    string       `json:"error"`
	UserInfo User         `json:"userInfo"`
	Friends  []FriendInfo `json:"friends"`
	Groups   []GroupInfo  `json:"groups"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterRes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	QQ     int `json:"qq`
	Status int `json:"status`
}
