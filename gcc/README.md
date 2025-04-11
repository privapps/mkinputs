# Cross-Platform Automation Tool

A C program that reads simple command files containing keyboard/mouse commands and executes them in sequence.

## Features
- Cross-platform support (macOS, Windows, Linux)
- Simple command format (no external dependencies)
- Full keyboard and mouse control

## Requirements
- **GCC** compiler
- **macOS**: CoreGraphics, CoreFoundation frameworks
- **Windows**: Win32 API
- **Linux**: X11, XTest libraries (`sudo apt-get install libxtst-dev` or similar)

## Installation

### macOS
```bash
gcc -c input_mac.c -o input_mac.o -framework CoreGraphics -framework CoreFoundation
gcc main.c input_mac.o -o automation_tool -framework CoreGraphics -framework CoreFoundation
```

### Windows
```bash
gcc -c input_win.c -o input_win.o
gcc main.c input_win.o -o automation_tool.exe -lgdi32 -luser32
```

### Linux
```bash
gcc -c input_linux.c -o input_linux.o -lX11 -lXtst
gcc main.c input_linux.o -o automation_tool -lX11 -lXtst
```

## Usage
```bash
./automation_tool commands.txt
```

## Simple Command Format Reference

Commands are specified one per line, with space-separated arguments.

### Key Commands
- `keydown <key1> <key2> ...`: Press one or more keys simultaneously.
  ```
  keydown shift a
  ```
  
- `keyup <key1> <key2> ...`: Release one or more keys.
  ```
  keyup shift a
  ```

### Mouse Commands
- `mousemove <x> <y>`: Move the mouse to absolute coordinates.
  ```
  mousemove 305 587
  ```

- `click`: Perform a left mouse click at the current position.
  ```
  click
  ```

### Timing
- `sleep <milliseconds>`: Pause execution for the specified duration.
  ```
  sleep 1000
  ```

## Supported Keys

### Basic Input
- **Letters**: `a` through `z`
- **Numbers**: `0` through `9`

### Special Keys
- `backspace`, `delete`, `enter`, `tab`, `esc`, `escape`
- Arrow keys: `up`, `down`, `left`, `right`
- Navigation: `home`, `end`, `pageup`, `pagedown`

### Function Keys
- F1-F12: `f1` through `f12` (F13-F20 may work on some platforms)

### Modifiers
- Command/Win: `cmd`, `lcmd`, `rcmd`
- Alt/Option: `alt`, `lalt`, `ralt`
- Control: `ctrl`, `lctrl`, `rctrl`, `control`
- Shift: `shift`, `lshift`, `rshift`
- `capslock`, `space`

### Media Controls (Platform Dependent)
- Volume: `audio_mute`, `audio_vol_down`, `audio_vol_up`
- Playback: `audio_play`, `audio_stop`, `audio_pause`
- Track: `audio_prev`, `audio_next`

### Numpad
- Numbers: `num0` through `num9`
- `num_lock`

## Example Command File (`commands.txt`)
```
keydown shift a
sleep 100
keyup shift a
mousemove 500 500
click
sleep 1000
keydown ctrl alt delete
sleep 50
keyup ctrl alt delete
```
