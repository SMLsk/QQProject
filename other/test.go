package main

import (
	"fmt"
	// "time"

	// "net/http"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	startWindow()
}

func startWindow() {
	fmt.Println("")
	Window, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{300, 100, 950, 630})
	if err != nil {
		// log.Fatal(err)
	}
	Window.LoadFile("E://Technology//Project//MyGo//src//QQ//application//client//view//dialog0Test.html")
	Window.SetTitle("Main")
	Window.Show()
	Window.Run()
}
