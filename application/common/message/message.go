package message

const (
	LoginMesType            = "LoginMes"
	LoginResType            = "LoginRes"
	RegisterMesType         = "RegisterMes"
	RegisterResType         = "RegisterRes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	ShortMessageMesType     = "ShortMessageMes"
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
	QQ     int `json:"qq"`
	Status int `json:"status"`
}

type ShortMessageMes struct {
	//私聊为0，群聊为1
	Type         int     `json:"type"`
	SenderId     int     `json:"senderId"`
	RecipientId  int     `json:"recipient"`
	Content      string  `json:"content"`
	Rows         float64 `json:"rows"`
	Length       float64 `json:"length"`
	ChartGroupId int     `json:"chartGroupId"`
	Time         int64   `json:"time"`
}
