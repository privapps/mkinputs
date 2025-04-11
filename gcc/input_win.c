#include "input.h"
#include <windows.h>

// Windows key codes
KeyMapping key_map[] = {
    {"a", 0x41}, {"b", 0x42}, {"c", 0x43}, {"d", 0x44}, {"e", 0x45},
    {"f", 0x46}, {"g", 0x47}, {"h", 0x48}, {"i", 0x49}, {"j", 0x4A},
    {"k", 0x4B}, {"l", 0x4C}, {"m", 0x4D}, {"n", 0x4E}, {"o", 0x4F},
    {"p", 0x50}, {"q", 0x51}, {"r", 0x52}, {"s", 0x53}, {"t", 0x54},
    {"u", 0x55}, {"v", 0x56}, {"w", 0x57}, {"x", 0x58}, {"y", 0x59},
    {"z", 0x5A},

    {"0", 0x30}, {"1", 0x31}, {"2", 0x32}, {"3", 0x33}, {"4", 0x34},
    {"5", 0x35}, {"6", 0x36}, {"7", 0x37}, {"8", 0x38}, {"9", 0x39},

    {"backspace", 0x08}, {"delete", 0x2E}, {"enter", 0x0D}, {"tab", 0x09},
    {"esc", 0x1B}, {"escape", 0x1B}, {"up", 0x26}, {"down", 0x28},
    {"right", 0x27}, {"left", 0x25}, {"home", 0x24}, {"end", 0x23},
    {"pageup", 0x21}, {"pagedown", 0x22},

    {"f1", 0x70}, {"f2", 0x71}, {"f3", 0x72}, {"f4", 0x73},
    {"f5", 0x74}, {"f6", 0x75}, {"f7", 0x76}, {"f8", 0x77},
    {"f9", 0x78}, {"f10", 0x79}, {"f11", 0x7A}, {"f12", 0x7B},

    {"cmd", 0x5B}, {"lcmd", 0x5B}, {"rcmd", 0x5C}, {"alt", 0x12},
    {"lalt", 0xA4}, {"ralt", 0xA5}, {"ctrl", 0x11}, {"lctrl", 0xA2},
    {"rctrl", 0x3E}, {"control", 0x11}, {"shift", 0x10}, {"lshift", 0xA0},
    {"rshift", 0xA1}, {"capslock", 0x14}, {"space", 0x20},

    {"audio_mute", 0xAD}, {"audio_vol_down", 0xAE}, {"audio_vol_up", 0xAF},
    {"audio_play", 0xB3}, {"audio_stop", 0xB2}, {"audio_pause", 0x13},
    {"audio_prev", 0xB1}, {"audio_next", 0xB0},

    {"num0", 0x60}, {"num1", 0x61}, {"num2", 0x62}, {"num3", 0x63},
    {"num4", 0x64}, {"num5", 0x65}, {"num6", 0x66}, {"num7", 0x67},
    {"num8", 0x68}, {"num9", 0x69}, {"num_lock", 0x90}
};

void simulate_key_event(KeyCode keycode, bool down) {
    INPUT input = {0};
    input.type = INPUT_KEYBOARD;
    input.ki.wVk = (WORD)keycode;
    input.ki.dwFlags = down ? 0 : KEYEVENTF_KEYUP;
    SendInput(1, &input, sizeof(INPUT));
}

void move_mouse(int x, int y) {
    INPUT input = {0};
    input.type = INPUT_MOUSE;
    input.mi.dwFlags = MOUSEEVENTF_ABSOLUTE | MOUSEEVENTF_MOVE;
    input.mi.dx = x * (65535.0f / (GetSystemMetrics(SM_CXSCREEN) - 1));
    input.mi.dy = y * (65535.0f / (GetSystemMetrics(SM_CYSCREEN) - 1));
    SendInput(1, &input, sizeof(INPUT));
}

void mouse_click() {
    INPUT input = {0};
    input.type = INPUT_MOUSE;
    input.mi.dwFlags = MOUSEEVENTF_LEFTDOWN;
    SendInput(1, &input, sizeof(INPUT));

    input.mi.dwFlags = MOUSEEVENTF_LEFTUP;
    SendInput(1, &input, sizeof(INPUT));
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
