package main

import (
	"fmt"
	"math"
	"unicode"
	// "time"
	// "net/http"
	// "github.com/sciter-sdk/go-sciter"
	// "github.com/sciter-sdk/go-sciter/window"
)

func main() {
	startWindow()
}

func startWindow() {
	// fmt.Println("")
	// Window, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{300, 100, 950, 630})
	// if err != nil {
	// 	// log.Fatal(err)
	// }
	// Window.LoadFile("E://Technology//Project//MyGo//src//QQ//application//client//view//dialog0Test.html")
	// Window.SetTitle("Main")
	// Window.Show()
	// Window.Run()
	str := "1234567890123456789012"
	var size float64
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			size += 2
		} else {
			size += 1
		}
	}
	rows := math.Ceil(size/24) + 1
	fmt.Println(size, rows)
}
