# quickstart для ncurses и argparse

import argparse
import curses
from curses import wrapper

HELP = """Press arrows
q  - exit
"""


def ncurses_tui(stdscr):  # noqa: CCR001
    stdscr.clear()
    curses.use_default_colors()
    stdscr.addstr(HELP)
    stdscr.refresh()
    global WP_NOW
    while True:
        c = stdscr.getch()
        stdscr.clear()
        stdscr.addstr(HELP)
        if c in [ord("q"), 208]:
            break

        if c == curses.KEY_LEFT:
            stdscr.addstr("key left\n")
        elif c == curses.KEY_RIGHT:
            stdscr.addstr("key right\n")
        elif c == curses.KEY_DOWN:
            stdscr.addstr("key down\n")
        elif c == curses.KEY_UP:
            stdscr.addstr("key up\n")


def main():
    parser = argparse.ArgumentParser(description="example of argparse and ncurses")
    parser.add_argument("--hello", action="store_true", help="returns world, overwrites other settings")
    parser.add_argument("-c", "--curses", action="store_true", help="default, returns interface with curses")
    args = parser.parse_args()

    if args.hello:
        print(r"\e[1;36mWorld")  # noqa: T201
    else:
        wrapper(ncurses_tui)


if __name__ == "__main__":
    main()
