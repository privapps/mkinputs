package main

import (
	"fmt"
	"time"

	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("--- ctrl + q to quit / ctrl + shift + hold to record position ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl"}, func(e hook.Event) {
		hook.End()
	})
	hook.Register(hook.MouseHold, []string{"ctrl", "shift"}, func(e hook.Event) {
		// x, y := robotgo.GetMousePos()
		fmt.Printf("%v mouse => %v, %v\n", time.Now().Format(time.RFC3339), e.X, e.Y)

	})

	s := hook.Start()
	<-hook.Process(s)
}
