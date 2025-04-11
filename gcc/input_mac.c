#include "input.h"
#include <CoreGraphics/CoreGraphics.h>

// macOS key codes
KeyMapping key_map[] = {
    // Letters
    {"a", 0x00}, {"b", 0x0B}, {"c", 0x08}, {"d", 0x02}, {"e", 0x0E},
    {"f", 0x03}, {"g", 0x05}, {"h", 0x04}, {"i", 0x22}, {"j", 0x26},
    {"k", 0x28}, {"l", 0x25}, {"m", 0x2E}, {"n", 0x2D}, {"o", 0x1F},
    {"p", 0x23}, {"q", 0x0C}, {"r", 0x0F}, {"s", 0x01}, {"t", 0x11},
    {"u", 0x20}, {"v", 0x09}, {"w", 0x0D}, {"x", 0x07}, {"y", 0x10},
    {"z", 0x06},

    // Numbers
    {"0", 0x1D}, {"1", 0x12}, {"2", 0x13}, {"3", 0x14}, {"4", 0x15},
    {"5", 0x17}, {"6", 0x16}, {"7", 0x1A}, {"8", 0x1C}, {"9", 0x19},

    // Special keys
    {"backspace", 0x33}, {"delete", 0x75}, {"enter", 0x24}, {"tab", 0x30},
    {"esc", 0x35}, {"escape", 0x35}, {"up", 0x7E}, {"down", 0x7D},
    {"right", 0x7C}, {"left", 0x7B}, {"home", 0x73}, {"end", 0x77},
    {"pageup", 0x74}, {"pagedown", 0x79},

    // Function keys
    {"f1", 0x7A}, {"f2", 0x78}, {"f3", 0x63}, {"f4", 0x76},
    {"f5", 0x60}, {"f6", 0x61}, {"f7", 0x62}, {"f8", 0x64},
    {"f9", 0x65}, {"f10", 0x6D}, {"f11", 0x67}, {"f12", 0x6F},
    {"f13", 0x69}, {"f14", 0x6B}, {"f15", 0x71}, {"f16", 0x6A},
    {"f17", 0x40}, {"f18", 0x4F}, {"f19", 0x50}, {"f20", 0x5A},

    // Modifiers
    {"cmd", 0x37}, {"lcmd", 0x37}, {"rcmd", 0x36}, {"alt", 0x3A},
    {"lalt", 0x3A}, {"ralt", 0x3D}, {"ctrl", 0x3B}, {"lctrl", 0x3B},
    {"rctrl", 0x3E}, {"control", 0x3B}, {"shift", 0x38}, {"lshift", 0x38},
    {"rshift", 0x3C}, {"capslock", 0x39}, {"space", 0x31},

    // Media keys (macOS supports these)
    {"audio_mute", 0x4A}, {"audio_vol_down", 0x49}, {"audio_vol_up", 0x48},
    {"audio_play", 0x34}, {"audio_stop", 0x66}, {"audio_pause", 0x71},
    {"audio_prev", 0x70}, {"audio_next", 0x69},

    // Numpad
    {"num0", 0x52}, {"num1", 0x53}, {"num2", 0x54}, {"num3", 0x55},
    {"num4", 0x56}, {"num5", 0x57}, {"num6", 0x58}, {"num7", 0x59},
    {"num8", 0x5B}, {"num9", 0x5C}, {"num_lock", 0x47}
};

void simulate_key_event(KeyCode keycode, bool down) {
    CGEventRef event = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)keycode, down);
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void move_mouse(int x, int y) {
    CGEventRef event = CGEventCreateMouseEvent(
        NULL, kCGEventMouseMoved,
        CGPointMake(x, y),
        kCGMouseButtonLeft
    );
    CGEventPost(kCGHIDEventTap, event);
    CFRelease(event);
}

void mouse_click() {
    // Left mouse down
    CGEventRef down = CGEventCreateMouseEvent(
        NULL, kCGEventLeftMouseDown,
        CGEventGetLocation(CGEventCreate(NULL)),
        kCGMouseButtonLeft
    );
    CGEventPost(kCGHIDEventTap, down);
    CFRelease(down);
    
    // Left mouse up
    CGEventRef up = CGEventCreateMouseEvent(
        NULL, kCGEventLeftMouseUp,
        CGEventGetLocation(CGEventCreate(NULL)),
        kCGMouseButtonLeft
    );
    CGEventPost(kCGHIDEventTap, up);
    CFRelease(up);
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
