#ifndef INPUT_H
#define INPUT_H

#ifdef _WIN32
  #include <windows.h>
#elif __linux__
  #include <X11/Xlib.h>
  #include <X11/extensions/XTest.h>
#else
  #include <CoreGraphics/CoreGraphics.h>
#endif

#ifndef KEY_MAPPING_H
#define KEY_MAPPING_H

// Key codes (platform-independent)
typedef int KeyCode;

typedef struct {
    const char *name;
    int code;
} KeyMapping;

extern KeyMapping key_map[];

#endif

void simulate_key_event(KeyCode keycode, bool down);
void move_mouse(int x, int y);
void mouse_click();

int get_key_code(const char *key_name);
void press_keys(char *keys);
void release_keys(char *keys);

#endif
