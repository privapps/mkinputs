## mkinputo: Automated Mouse & Keyboard Input via YAML

### What
This project provides small, cross-platform Go programs to automate mouse and keyboard interactions, all described in a simple YAML file.

---

## Usage
### mkinput (main executor)
**Run actions from a YAML file:**
```bash
./mkinput actions.yaml [extra arguments]
```
- **actions.yaml**: Required. A YAML automation sequence (see below).
- **[extra arguments]**: Optional. Used for `typearg` actions in YAML.
- `--verbose`/`-v`: Optional. Print each action as it is executed (for debugging).

**!! Security Warning: DO NOT store real passwords or secrets in plain YAML: use environment variables.**

**Example:**
```bash
export MY_PASS=supersecret
./mkinput reboot.yaml my_server_username
```
And in YAML:
```yaml
- typestr: ${MY_PASS}
```
All `${VARNAME}` expressions are replaced by their current environment variable.

#### Supported YAML Actions
| Action      | Example/Args                           | Effect                                                        |
|-------------|----------------------------------------|---------------------------------------------------------------|
| mouse       | `- mouse: x y`                         | Move mouse to absolute position (x, y)                        |
| move        | `- move: dx dy`                        | Move mouse relative by (dx, dy)                               |
| click       | `- click:` or `- click: left`          | Mouse click, optionally specifying button                     |
| drag        | `- drag: x y`                          | Hold mouse and drag to (x, y)                                 |
| sleep       | `- sleep: ms`                          | Wait for ms milliseconds                                      |
| keytab      | `- keytab: enter shift`                | Press tab/key combination                                     |
| keydown     | `- keydown: shift`                     | Hold down a key                                               |
| keyup       | `- keyup: shift`                       | Release a key                                                 |
| typestr     | `- typestr: your text here`            | Type a string of text                                         |
| typearg     | `- typearg: n`                         | Type CLI argument n (os.Args[n])                              |
| keytoggle   | `- keytoggle: key [down|up]`           | Toggle key on/off                                             |
| toggle      | `- toggle: key kind`                   | For specialty toggles                                         |
| typespace   | `- typespace: n`                       | Type n spaces                                                 |

  For details and key codes, see: [robotgo keys](https://github.com/go-vgo/robotgo/blob/master/docs/keys.md)

### track (mouse position recorder)
Run with no arguments:
```bash
./track
```
**While running:**
- `ctrl+q` — Quit
- `ctrl+shift` + hold mouse — Print current mouse coordinates (for YAML editing)

---

## Example YAML Scenarios

### 1. Make Mac Sleep via Mouse (1450x900 screen)
```yaml
- mouse: 29 12
- click:
- mouse: 29 196
- click:
```

### 2. Login via Keyboard (Any OS)
```yaml
- keytab: tab alt
- keyup: shift
- sleep: 100
- typestr: <username>
- keytab: tab
- typestr: ${LOGIN_PASS}
- keytab: tab
- keytab: enter
```

### 3. SSH & Reboot multiple servers:
```bash
for i in ip1 ip2 ip3; do ./mkinput reboot.yaml $i; done
```
```yaml
- typestr: ssh <username>@
- typearg: 2   # Use command-line arg for hostname/user
- keytab: enter
- sleep: 2000
- typestr: ${SSH_PASS}
- keytab: enter
- sleep: 1000
- typestr: sudo reboot
- keytab: enter
- sleep: 800
- typestr: ${SSH_PASS}
- keytab: enter
- sleep: 1000
```

---

## Command-line Reference

### mkinput
- Usage: `./mkinput action.yaml [args...]`
- **Args:**
  - `action.yaml`: Path to required YAML config
  - `[args...]`: Extra args (referenced via `typearg`)
  - `--verbose` or `-v`: Print each action command as it's executed.
- **No other flags**. Invalid/missing YAML results in a clear error and (if present) line number.

### track
- Usage: `./track`
- No arguments or flags

---

## Troubleshooting
- **Missing YAML**: mkinput will exit with "Yaml file required".
- **Parse errors**: The error message now shows the line number and a sample of the faulty line.
- **Windows slowness**: disabling antivirus may help.
- **OS-specific features**: Not all actions are supported on every OS.
- ctrl+q or ctrl+shift (track) not working? Ensure the window has focus and input method is English/US.

---

## Supported OSes
- macOS (all modern versions)
- Linux (most x86/arm distros)
- Windows (via Cygwin or similar, caveats apply)

---

## Linting

This project uses golangci-lint to ensure code quality. Check or install via:
```bash
golangci-lint run
```
See `.golangci.yml` for configuration. Lint runs automatically in CI.

---

## Additional Notes
- See [robotgo](https://github.com/go-vgo/robotgo) for low-level input docs/limitations.
- Any questions, issues, or feature requests: please file on GitHub.

---