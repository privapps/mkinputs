#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "input.h"

#ifdef _WIN32
//#include "input_win.c"
#elif __linux__
//#include "input_linux.c"
#else
//#include "input_mac.c"
#endif

typedef enum {
    CMD_KEYDOWN,
    CMD_KEYUP,
    CMD_MOUSEMOVE,
    CMD_CLICK,
    CMD_SLEEP
} CommandType;

typedef struct {
    CommandType type;
    union {
        struct {
            char *keys; // For keydown/keyup
        };
        struct {
            int x, y;   // For mousemove
        };
        int duration;   // For sleep
    };
} Command;

int main(int argc, char **argv) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <command_file>\n", argv[0]);
        return 1;
    }

    FILE *file = fopen(argv[1], "r");
    if (!file) {
        fprintf(stderr, "Failed to open file: %s\n", argv[1]);
        return 1;
    }

    char *line = NULL;
    size_t len = 0;
    ssize_t read;

    while ((read = getline(&line, &len, file)) != -1) {
        // Remove newline character
        if (line[read - 1] == '\n') {
            line[read - 1] = '\0';
        }

        char *command = strtok(line, " ");
        if (command == NULL) continue;

        if (strcmp(command, "keydown") == 0) {
            char *keys = strtok(NULL, "");
            if (keys != NULL) {
                press_keys(keys);
            }
        } else if (strcmp(command, "keyup") == 0) {
            char *keys = strtok(NULL, "");
            if (keys != NULL) {
                release_keys(keys);
            }
        } else if (strcmp(command, "mousemove") == 0) {
            char *x_str = strtok(NULL, " ");
            char *y_str = strtok(NULL, " ");
            if (x_str != NULL && y_str != NULL) {
                int x = atoi(x_str);
                int y = atoi(y_str);
                move_mouse(x, y);
            }
        } else if (strcmp(command, "click") == 0) {
            mouse_click();
        } else if (strcmp(command, "sleep") == 0) {
            char *duration_str = strtok(NULL, " ");
            if (duration_str != NULL) {
                int duration = atoi(duration_str);
                usleep(duration * 1000);
            }
        }
    }

    fclose(file);
    if (line) free(line);

    return 0;
}
