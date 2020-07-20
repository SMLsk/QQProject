# 完成上线下线通知

## 一、完成上线通知：

1、构建Message

```go
const NotifyUserStatusMesType = "NotifyUserStatusMes"
type NotifyUserStatusMes struct {
	QQ     int `json:"qq`
	Status int `json:"status`
}
```

2、完成通知函数

(1)给UserProcessor添加friends字段,在登录后将好友表传入

```go
type UserProcessor struct {
	Conn    net.Conn
	QQ      int
	Friends []int
}
```

```go
func (this *UserProcess) NotifyOtherOnlineUsers(staus int) (err error) {
    for _,qq := range this.Friends{
        up,err = MyUserManager.GetOnlineUserByQQ(qq)
        if err == nil{
            err = up.NotifyMeUserStatus(this.QQ,status)
            if err != nil {
			return
			}
        }
    }
	return
}

func (this *UserProcess) NotifyMeUserStatus(qq int,status common.Online) (err error) {
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = status
    
	data, err := common.AssembleMes(message.NotifyUserStatusMesType, notifyUserStatusMes)
	if err != nil {
		return
	}

	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)

	return
}
```

## Clinet接收消息







