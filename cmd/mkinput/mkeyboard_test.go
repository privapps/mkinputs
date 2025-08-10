package main

import (
	"os"
	"testing"
)

// MockRobotActions for test safety
// All methods implement RobotActions but do nothing.
type MockRobotActions struct{}

func (m *MockRobotActions) KeyTap(key string, args ...interface{})    {}
func (m *MockRobotActions) KeyUp(key string)                          {}
func (m *MockRobotActions) KeyDown(key string)                        {}
func (m *MockRobotActions) TypeStr(s string)                          {}
func (m *MockRobotActions) KeyToggle(key string, args ...interface{}) {}
func (m *MockRobotActions) MilliSleep(ms int)                         {}
func (m *MockRobotActions) MoveSmooth(x, y int, args ...interface{})  {}
func (m *MockRobotActions) MoveRelative(x, y int)                     {}
func (m *MockRobotActions) Click(args ...interface{})                 {}
func (m *MockRobotActions) DragSmooth(x, y int, args ...interface{})  {}
func (m *MockRobotActions) Toggle(key, kind string)                   {}

func TestGetCmds(t *testing.T) {
	// Given
	yamlData := "---\n- typestr: ssh <username>@\n- keytab: enter\n"
	err := os.WriteFile("test.yaml", []byte(yamlData), 0644)
	if err != nil {
		t.Fatalf("Failed to write yaml file: %v", err)
	}

	// Patch os.Args for test
	oldArgs := os.Args
	os.Args = []string{"cmd", "test.yaml"}
	defer func() { os.Args = oldArgs }()

	// When
	cmds := getCmds()

	// Then
	if len(cmds) == 0 {
		t.Error("Expected commands but got none")
	}
	os.Remove("test.yaml")
}

func TestAct(t *testing.T) {
	// Given
	actions := &MockRobotActions{}
	cmds := []CEntry{{cmd: "typestr", args: "Hello World"}, {cmd: "sleep", args: "1000"}}
	// When
	act(actions, cmds)
	// Then check if function executes without panic
}

func TestActAllActions(t *testing.T) {
	tests := []struct {
		name    string
		cmds    []CEntry
		errorOk bool
	}{
		{"keytab", []CEntry{{cmd: "keytab", args: "enter"}}, true},
		{"keyup", []CEntry{{cmd: "keyup", args: "shift"}}, true},
		{"keydown", []CEntry{{cmd: "keydown", args: "shift"}}, true},
		{"typestr", []CEntry{{cmd: "typestr", args: "hello"}}, true},
		{"keytoggle-1", []CEntry{{cmd: "keytoggle", args: "a"}}, true},
		{"keytoggle-2", []CEntry{{cmd: "keytoggle", args: "a up"}}, true},
		{"sleep", []CEntry{{cmd: "sleep", args: "1"}}, true},
		{"mouse", []CEntry{{cmd: "mouse", args: "100 200"}}, true},
		{"mouse-more", []CEntry{{cmd: "mouse", args: "100 200 0.2 0.2"}}, true},
		{"move", []CEntry{{cmd: "move", args: "10 20"}}, true},
		{"click-0", []CEntry{{cmd: "click", args: ""}}, true},
		{"click-1", []CEntry{{cmd: "click", args: "right"}}, true},
		{"click-2", []CEntry{{cmd: "click", args: "right true"}}, true},
		{"drag", []CEntry{{cmd: "drag", args: "10 10"}}, true},
		{"drag-more", []CEntry{{cmd: "drag", args: "10 10 0.1 0.1"}}, true},
		{"toggle", []CEntry{{cmd: "toggle", args: "left up"}}, true},
		{"typespace", []CEntry{{cmd: "typespace", args: "2"}}, true},
		{"typearg", []CEntry{{cmd: "typearg", args: "1"}}, true},
		// error conditions
		{"bad-mouse-args", []CEntry{{cmd: "mouse", args: "100"}}, false},
		{"bad-move-args", []CEntry{{cmd: "move", args: "10"}}, false},
		{"bad-drag-args", []CEntry{{cmd: "drag", args: "10"}}, false},
		{"bad-toggle-args", []CEntry{{cmd: "toggle", args: ""}}, false},
		{"bad-typearg-args", []CEntry{{cmd: "typearg", args: ""}}, false},
		{"bad-keytoggle", []CEntry{{cmd: "keytoggle", args: ""}}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actions := &MockRobotActions{}
			defer func() {
				if r := recover(); r != nil {
					if tc.errorOk {
						t.Errorf("action %s: unexpected panic: %v", tc.name, r)
					}
				} else {
					if !tc.errorOk {
						t.Errorf("action %s: expected panic, got none", tc.name)
					}
				}
			}()
			act(actions, tc.cmds)
		})
	}
}
