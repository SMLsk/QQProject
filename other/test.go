package main

import (
	"fmt"
	// "github.com/sciter-sdk/go-sciter"
	// "github.com/sciter-sdk/go-sciter/window"
)

type FriendsManager struct {
	Groups        map[int]string
	GroupsFriends map[int]int
}

var (
	MyFriendsManager *FriendsManager
)

func main() {
	fmt.Println(MyFriendsManager)
}
