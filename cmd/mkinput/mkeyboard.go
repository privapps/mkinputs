package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"io/ioutil"
	"strconv"

	"github.com/go-vgo/robotgo"
	"gopkg.in/yaml.v3"
)

// Entry represents a mapping of keys to their string values.
type Entry map[string]string

// CEntry represents a command entry with its arguments.
type CEntry struct {
	cmd  string
	args string
}

// getCmds loads a YAML file (from os.Args[1]) as a sequence of actions/commands.
// Extra CLI args are available for 'typearg' injection in YAML.
// getCmds loads a YAML file (from os.Args[1]) as a sequence of actions/commands.
// Exits with error+line if YAML invalid.
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
		// Enhanced error handling: reparse as yaml.Node for line number,
		// print a pretty parse error with context
		var nodes []yaml.Node
		err3 := yaml.Unmarshal(source, &nodes)
		if err3 == nil && len(nodes) > 0 {
			fmt.Printf("YAML parse error at line %d, column %d: %v\nSample line: %s\n", nodes[0].Line, nodes[0].Column, err2, linesAt(nodes[0].Line, source))
		} else {
			fmt.Printf("Failed to parse YAML file: %v\n", err2)
		}
		os.Exit(1)
	}
	var cmds []CEntry
	for _, v := range data {
		for k2, v2 := range v {
			repl := os.Expand(strings.TrimSpace(v2), func(env string) string { return os.Getenv(env) })
			cmds = append(cmds, CEntry{cmd: strings.TrimSpace(k2), args: repl})
		}
	}
	return cmds
}

// linesAt returns the content of the given (1-based) line number from src.
func linesAt(line int, src []byte) string {
	if line < 1 {
		return "(unknown line)"
	}
	lines := strings.Split(string(src), "\n")
	if line > len(lines) {
		return "(line out of bounds)"
	}
	return lines[line-1]
}

// RobotActions defines the set of methods needed to perform input actions (mouse/keyboard).
// Used to abstract and mock robotgo for testing or alternate backends.
type RobotActions interface {
	// KeyTap simulates pressing specified key(s) once.
	KeyTap(key string, args ...interface{})
	// KeyUp releases a held-down key.
	KeyUp(key string)
	// KeyDown presses a key down (no release).
	KeyDown(key string)
	// TypeStr types a full string using the keyboard.
	TypeStr(s string)
	// KeyToggle presses or releases a key, possibly specifying up/down state.
	KeyToggle(key string, args ...interface{})
	// MilliSleep sleeps for given milliseconds.
	MilliSleep(ms int)
	// MoveSmooth moves mouse smoothly to (x, y). Optional speed/curve args.
	MoveSmooth(x, y int, args ...interface{})
	// MoveRelative moves mouse by (x, y) relative to current position.
	MoveRelative(x, y int)
	// Click performs a mouse click, with optional button and bool args.
	Click(args ...interface{})
	// DragSmooth holds and drags the mouse to (x, y).
	DragSmooth(x, y int, args ...interface{})
	// Toggle is for specialty (nonstandard) key toggles.
	Toggle(key, kind string)
}

// RobotGoActions implements RobotActions with the real robotgo backend.
type RobotGoActions struct{}

func (r *RobotGoActions) KeyTap(key string, args ...interface{}) {
	robotgo.KeyTap(key, args...)
}
func (r *RobotGoActions) KeyUp(key string)   { robotgo.KeyUp(key) }
func (r *RobotGoActions) KeyDown(key string) { robotgo.KeyDown(key) }
func (r *RobotGoActions) TypeStr(s string)   { robotgo.TypeStr(s) }
func (r *RobotGoActions) KeyToggle(key string, args ...interface{}) {
	robotgo.KeyToggle(key, args...)
}
func (r *RobotGoActions) MilliSleep(ms int)                        { robotgo.MilliSleep(ms) }
func (r *RobotGoActions) MoveSmooth(x, y int, args ...interface{}) { robotgo.MoveSmooth(x, y, args...) }
func (r *RobotGoActions) MoveRelative(x, y int)                    { robotgo.MoveRelative(x, y) }
func (r *RobotGoActions) Click(args ...interface{})                { robotgo.Click(args...) }
func (r *RobotGoActions) DragSmooth(x, y int, args ...interface{}) { robotgo.DragSmooth(x, y, args...) }
func (r *RobotGoActions) Toggle(key, kind string)                  { robotgo.Toggle(key, kind) }

// act executes a list of action commands using the provided RobotActions handler.
// This is the main engine for YAML-driven input automation.
// act executes a list of action commands using the provided RobotActions handler.
// This is the main engine for YAML-driven input automation.
// act runs the YAML actions using the RobotActions backend. Logs if verbose mode.
type actionFunc func(actions RobotActions, args []string, cmd CEntry)

// actionRegistry maps command names to implementation functions
var actionRegistry = map[string]actionFunc{
	"keytab": func(a RobotActions, args []string, cmd CEntry) {
		if len(args) == 1 {
			a.KeyTap(args[0])
		} else {
			ifaceArgs := make([]interface{}, len(args[1:]))
			for i, v := range args[1:] {
				ifaceArgs[i] = v
			}
			a.KeyTap(args[0], ifaceArgs...)
		}
	},
	"keyup":   func(a RobotActions, args []string, _ CEntry) { a.KeyUp(args[0]) },
	"keydown": func(a RobotActions, args []string, _ CEntry) { a.KeyDown(args[0]) },
	"typestr": func(a RobotActions, args []string, cmd CEntry) { a.TypeStr(cmd.args) },
	"keytoggle": func(a RobotActions, args []string, cmd CEntry) {
		if len(args) < 1 {
			panic("key toggle missing args")
		}
		switch len(args) {
		case 1:
			a.KeyToggle(args[0])
		default:
			ifaceArgs := make([]interface{}, len(args[1:]))
			for i, v := range args[1:] {
				ifaceArgs[i] = v
			}
			a.KeyToggle(args[0], ifaceArgs...)
		}
	},
	"sleep": func(a RobotActions, args []string, _ CEntry) {
		s, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		a.MilliSleep(s)
	},
	"mouse": func(a RobotActions, args []string, _ CEntry) {
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
			ifaceArgs := make([]interface{}, len(args[2:]))
			for idx, val := range args[2:] {
				ifaceArgs[idx] = val
			}
			a.MoveSmooth(x, y, ifaceArgs...)
		} else {
			a.MoveSmooth(x, y, 0.3, 0.3)
		}
	},
	"move": func(a RobotActions, args []string, _ CEntry) {
		if len(args) != 2 {
			panic("mouse move requires x, y")
		}
		x, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		a.MoveRelative(x, y)
	},
	"click": func(a RobotActions, args []string, _ CEntry) {
		switch len(args) {
		case 0:
			a.Click()
		case 1:
			a.Click(args[0])
		case 2:
			b, err := strconv.ParseBool(args[1])
			if err != nil {
				panic(err)
			}
			a.Click(args[0], b)
		}
	},
	"drag": func(a RobotActions, args []string, _ CEntry) {
		if len(args) < 2 {
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
			ifaceArgs := make([]interface{}, len(args[2:]))
			for idx, val := range args[2:] {
				ifaceArgs[idx] = val
			}
			a.DragSmooth(x, y, ifaceArgs...)
		} else {
			a.DragSmooth(x, y, 0.3, 0.3)
		}
	},
	"toggle": func(a RobotActions, args []string, _ CEntry) {
		if len(args) < 1 {
			panic("missing args")
		}
		a.Toggle(args[0], args[1])
	},
	"typespace": func(a RobotActions, args []string, _ CEntry) {
		sc := 1
		if len(args) > 0 {
			arg2, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			sc = arg2
		}
		for i := 0; i < sc; i++ {
			a.TypeStr(" ")
		}
	},
	"typearg": func(a RobotActions, args []string, _ CEntry) {
		if len(args) != 1 {
			panic("only one argument expected")
		}
		position, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		a.TypeStr(os.Args[position])
	},
}

func act(actions RobotActions, cmds []CEntry) {
	for i, cmd := range cmds {
		args := strings.Fields(cmd.args)
		if fn, ok := actionRegistry[strings.ToLower(cmd.cmd)]; ok {
			logVerbose("%d: %s %q\n", i, cmd.cmd, cmd.args)
			fn(actions, args, cmd)
		} else {
			fmt.Printf("Unknown cmd: %s\n", cmd.cmd)
		}
	}
}

var verboseMode = false

// main is the entrypoint for mkinput. Parses CLI, runs YAML actions.

func main() {
	// Cross-platform support guard
	supported := map[string]bool{"darwin": true, "linux": true, "windows": true}
	if !supported[runtime.GOOS] {
		fmt.Printf("Warning: OS '%s' is untested and may not be supported.\n", runtime.GOOS)
	}
	// Simple verbose flag check
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--verbose") {
		verboseMode = true
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
	cmds := getCmds()
	actions := &RobotGoActions{}
	act(actions, cmds)
}

// logVerbose prints to stdout if --verbose enabled.
func logVerbose(format string, args ...interface{}) {
	if verboseMode {
		fmt.Printf(format, args...)
	}
}
