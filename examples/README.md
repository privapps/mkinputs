# mkinputo YAML Examples Library

This directory contains real-world YAML workflow examples for mkinputo.

---

## sleep-mac.yaml
Automate clicks to make a Mac go to sleep.
**Expected Output:**
- Mouse will move to (29, 12) and click (menu)
- Mouse will move to (29, 196) and click (sleep)

---

## ssh-reboot.yaml
Login and reboot remote hosts via SSH using password from env var.
**Expected Output:**
- Types and logs into servers, sends password, types `sudo reboot`, and confirms with password again.
- Can be run in a loop over multiple servers.

---

Custom YAMLs can be created easily: just refer to this directory for syntax.

For more, see mkinputo main README and [robotgo key codes](https://github.com/go-vgo/robotgo/blob/master/docs/keys.md).
