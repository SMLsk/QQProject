package utils

import (
	"QQ/application/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		err = io.EOF
		return
	}
	dataLen := binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:dataLen])
	if uint32(n) != dataLen || err != nil {
		//数据接收失败
		return
	}
	//反序列化
	//这里可能出问题，mes是需要传入一个指针，虽然已经将mes定义为了指针
	//但是不确定有没有其他问题
	err = json.Unmarshal(this.Buf[:dataLen], &mes)
	if err != nil {
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {

	dataLen := uint32(len(data))
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, dataLen)
	n, err := this.Conn.Write(buf)
	if n != 4 || err != nil {
		fmt.Println("长度发送失败")
		return
	}

	n, err = this.Conn.Write(data)
	if uint32(n) != dataLen || err != nil {
		err = errors.New("data发送失败")
		return
	}
	return
}
