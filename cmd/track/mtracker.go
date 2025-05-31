package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("--- ctrl + q to quit / ctrl + shift + click to record mouse position ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl"}, func(e hook.Event) {
		hook.End()
	})
	hook.Register(hook.MouseDown, []string{"ctrl", "shift"}, func(e hook.Event) {
		x, y := robotgo.Location()
		color := robotgo.GetPixelColor(x-2, y-2) // must shift a little bit as it will color of mouse point.
		fmt.Printf("%v mouse => %v\t%v\t%v\n", time.Now().Format(time.RFC3339), x, y, color)
	})

	s := hook.Start()
	<-hook.Process(s)
}
