package main

import (
	"fmt"
	"os"
	"strings"

	"io/ioutil"
	"strconv"

	"github.com/go-vgo/robotgo"
	"gopkg.in/yaml.v3"
)

type Entry map[string]string
type CEntry struct {
	cmd  string
	args string
}

func getCmds() []CEntry {
	if len(os.Args) <= 1 {
		fmt.Println("Yaml file required")
		os.Exit(1)
	}
	file := os.Args[1]
	source, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Failed reading config file: %v, %v\n", file, err)
		os.Exit(1)
	}
	var data []Entry

	err2 := yaml.Unmarshal(source, &data)
	if err2 != nil {
		panic(err2)
	}
	var cmds []CEntry
	for _, v := range data {
		for k2, v2 := range v {
			cmds = append(cmds, CEntry{cmd: strings.TrimSpace(k2), args: strings.TrimSpace(v2)})
		}
	}
	return cmds
}

func act(cmds []CEntry) {
	for i, cmd := range cmds {
		args := strings.Fields(cmd.args)
		switch strings.ToLower(cmd.cmd) {
		case "keytab":
			if len(args) == 1 {
				robotgo.KeyTap(args[0])
			} else {
				robotgo.KeyTap(args[0], args[1:])
			}
		case "keyup":
			robotgo.KeyUp(args[0])
		case "keydown":
			robotgo.KeyDown(args[0])
		case "typestr":
			robotgo.TypeStr(strings.Join(args, " "))
		case "keytoggle":
			if len(args) < 1 {
				panic("key toggle missing args")
			}
			switch len(args) {
			case 1:
				robotgo.KeyToggle(args[0])
			default:
				robotgo.KeyToggle(args[0], args[1:]...)
			}
		case "sleep":
			s, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			robotgo.MilliSleep(s)
		case "mouse":
			if len(args) < 2 {
				panic("mouse requires x, y")
			}
			x, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			if len(args) > 2 {
				robotgo.MoveSmooth(x, y, args[2:])
			} else {
				robotgo.MoveSmooth(x, y, 0.3, 0.3)
			}
		case "click":
			switch len(args) {
			case 0:
				robotgo.Click()
			case 1:
				robotgo.Click(args[0])
			case 2:
				b, err := strconv.ParseBool(args[1])
				if err != nil {
					panic(err)
				}
				robotgo.Click(args[0], b)
			}
		case "drag":
			if len(args) != 2 {
				panic("mouse drag requires x, y")
			}
			x, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			if len(args) > 2 {
				robotgo.DragSmooth(x, y, args[2:])
			} else {
				robotgo.DragSmooth(x, y, 0.3, 0.3)
			}
		case "toggle":
			if len(args) < 1 {
				panic("missing args")
			}
			robotgo.Toggle(args...)
		case "typespace":
			var sc = 1
			if len(args) > 0 {
				arg2, err := strconv.Atoi(args[0])
				if err != nil {
					panic(err)
				}
				sc = arg2
			}
			for i := 0; i < sc; i++ {
				robotgo.TypeStr(" ")
			}
		case "typearg":
			if len(args) != 1 {
				panic("only one argument expected")
			}
			position, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			robotgo.TypeStr(os.Args[position])
		default:
			fmt.Println("unknown cmd " + cmd.cmd)
		}
		fmt.Printf("%d => %v\n", i, cmd.cmd)
	}
}

func main() {
	cmds := getCmds()
	act(cmds)
}
