#include "input.h"
#include <X11/Xlib.h>
#include <X11/extensions/XTest.h>

// Linux key codes (X11 key codes)
KeyMapping key_map[] = {
    {"a", 0x26}, {"b", 0x2c}, {"c", 0x2e}, {"d", 0x28}, {"e", 0x1a},
    {"f", 0x29}, {"g", 0x2a}, {"h", 0x2b}, {"i", 0x33}, {"j", 0x34},
    {"k", 0x35}, {"l", 0x36}, {"m", 0x38}, {"n", 0x37}, {"o", 0x39},
    {"p", 0x3a}, {"q", 0x18}, {"r", 0x1c}, {"s", 0x17}, {"t", 0x19},
    {"u", 0x30}, {"v", 0x2f}, {"w", 0x16}, {"x", 0x2d}, {"y", 0x1b},
    {"z", 0x25},

    {"0", 0x0b}, {"1", 0x02}, {"2", 0x03}, {"3", 0x04}, {"4", 0x05},
    {"5", 0x06}, {"6", 0x07}, {"7", 0x08}, {"8", 0x09}, {"9", 0x0a},

    {"backspace", 0x16}, {"delete", 0x71}, {"enter", 0x24}, {"tab", 0x0f},
    {"esc", 0x09}, {"escape", 0x09}, {"up", 0x6f}, {"down", 0x74},
    {"right", 0x72}, {"left", 0x71}, {"home", 0x6e}, {"end", 0x77},
    {"pageup", 0x70}, {"pagedown", 0x75},

    {"cmd", 0x64}, {"lcmd", 0x64}, {"rcmd", 0x64}, {"alt", 0x40},
    {"lalt", 0x40}, {"ralt", 0x40}, {"ctrl", 0x37}, {"lctrl", 0x37},
    {"rctrl", 0x37}, {"control", 0x37}, {"shift", 0x32}, {"lshift", 0x32},
    {"rshift", 0x32}, {"capslock", 0x48}, {"space", 0x41},

    {"audio_mute", 0x80}, {"audio_vol_down", 0x79}, {"audio_vol_up", 0x81},
    {"audio_play", 0xa2}, {"audio_stop", 0xa3}, {"audio_pause", 0xa4},
    {"audio_prev", 0x98}, {"audio_next", 0x99},

    {"num0", 0x52}, {"num1", 0x53}, {"num2", 0x54}, {"num3", 0x55},
    {"num4", 0x56}, {"num5", 0x57}, {"num6", 0x58}, {"num7", 0x59},
    {"num8", 0x5b}, {"num9", 0x5c}, {"num_lock", 0x4d}
};

static Display *display = NULL;

void simulate_key_event(KeyCode keycode, bool down) {
    if (!display) {
        display = XOpenDisplay(NULL);
    }
    XTestFakeKeyEvent(display, keycode, down, CurrentTime);
    XSync(display, False);
}

void move_mouse(int x, int y) {
    if (!display) {
        display = XOpenDisplay(NULL);
    }
    XTestFakeMotionEvent(display, 0, x, y, CurrentTime);
    XSync(display, False);
}

void mouse_click() {
    if (!display) {
        display = XOpenDisplay(NULL);
    }
    XTestFakeButtonEvent(display, Button1, True, CurrentTime);
    XTestFakeButtonEvent(display, Button1, False, CurrentTime);
    XSync(display, False);
}

int get_key_code(const char *key_name) {
    for (size_t i = 0; i < sizeof(key_map)/sizeof(KeyMapping); i++) {
        if (strcmp(key_name, key_map[i].name) == 0) {
            return key_map[i].code;
        }
    }
    return -1; // Invalid key code
}

void press_keys(char *keys) {
    char *token = strtok(keys, " ");
    while (token != NULL) {
        int code = get_key_code(token);
        if (code != -1) {
            simulate_key_event(code, true);
        }
        token = strtok(NULL, " ");
    }
}

void release_keys(char *keys) {
    char *token = strtok(keys, " ");
    while (token != NULL) {
        int code = get_key_code(token);
        if (code != -1) {
            simulate_key_event(code, false);
        }
        token = strtok(NULL, " ");
    }
}
