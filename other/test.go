package main

import (
	// "fmt"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	Window, err := window.New(sciter.DefaultWindowCreateFlag, &sciter.Rect{1190, 70, 1480, 630})
	if err != nil {
		// log.Fatal(err)
	}
	Window.LoadFile("E://Technology//Project//MyGo//src//QQ//application//client//view//mainTest.html")
	Window.SetTitle("Main")
	root, _ := Window.GetRootElement()
	setWinHandler(Window, root)
	Window.Show()
	Window.Run()
}
func setWinHandler(w *window.Window, root *sciter.Element) {
	w.DefineFunction("GetFriendsData", func(args ...*sciter.Value) *sciter.Value {

		return sciter.NewValue("{'a':'b'}")
	})
}
