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
	Window, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{1190, 70, 1480, 630})
	if err != nil {
		// log.Fatal(err)
	}
	Window.LoadFile("./test.html")
	Window.SetTitle("Main")
	// root, _ := Window.GetRootElement()
	Window.DefineFunction("startNewWindow", func(args ...*sciter.Value) *sciter.Value {
		Window2, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{1190, 70, 1480, 630})
		if err != nil {
			fmt.Println(err)
		}
		Window2.LoadFile("./test.html")
		Window2.SetTitle("Main")
		Window2.Show()
		Window2.Run()
		return sciter.NewValue("0")
	})
	Window.Show()
	Window.Run()
}
