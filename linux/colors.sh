#!/bin/bash

BLC='\e[0;30m'
RED='\e[1;31m'
GRN='\e[2;32m'
YEL='\e[3;33m'
BLU='\e[4;34m'
MAG='\e[5;35m'
CYA='\e[6;36m'
ENDC='\e[m'

# Формат записи цвета:
# "\e[<код>;<номер цвета>m"
# код:
# 0 - reset
# 1 - bold
# 2 - faint
# 3 - italic
# 4 - underline
# 5 - blinking
# 6 -
# more: https://en.wikipedia.org/wiki/ANSI_escape_code#SGR
# or here: http://www.andrewnoske.com/wiki/Bash_-_adding_color

printf "${BLC}black normal${ENDC}\n"
printf "${RED}red bold${ENDC}\n"
printf "${GRN}green faint(like normal)${ENDC}\n"
printf "${YEL}yellow italic(or background)${ENDC}\n"
printf "${BLU}blue underline${ENDC}\n"
printf "${MAG}magenta blinking${ENDC}\n"
printf "${CYA}cyan blinking(or normal)${ENDC}\n"
printf "white normal\n"

# В питоне формат иной(вместо "\e" писать "\033"), например:
#
# RED   = "\033[91m"
# YELLOW = "\033[93m"
# BLUE  = "\033[94m"
# CYAN  = "\033[36;4m"
# MAGENTA  = "\033[35;1m"
# GREEN = "\033[92m"
# ENDC = '\033[0m'
# print(f'{CYAN}Linked{ENDC}')
