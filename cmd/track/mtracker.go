package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println(startupMessage())
	hook.Register(hook.KeyDown, []string{"q", "ctrl"}, func(e hook.Event) {
		hook.End()
	})
	hook.Register(hook.MouseHold, []string{"ctrl", "shift"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		fmt.Print(formatMouseEvent(time.Now(), x, y))
	})

	s := hook.Start()
	<-hook.Process(s)
}

func formatMouseEvent(t time.Time, x, y int) string {
	return fmt.Sprintf("%v mouse => %v, %v\n", t.Format(time.RFC3339), x, y)
}

func startupMessage() string {
	return "--- ctrl + q to quit / ctrl + shift + hold to record position ---"
}
