#!/bin/bash
# Скрипт установки софта на новый сервер
apt update
apt upgrade -y

apt install -y python3 python3-pip python3-dev python3-setuptools python3-venv htop neofetch tmux zsh git curl wget vim build-essential

cd /tmp
git clone https://gitlab.com/teadove/dotfiles
cd dotfiles
./dotfiles_setup.py -r

chsh -s $(which zsh)
