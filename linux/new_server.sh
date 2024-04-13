#!/bin/bash
# Скрипт установки софта на новый сервер
apt update
apt upgrade -y

apt install -y python3 python3-pip python3-dev python3-setuptools python3-venv htop neofetch tmux zsh git curl wget vim build-essential net-tools make

cd /tmp
git clone https://gitlab.com/teadove/dotfiles
cd dotfiles
./dotfiles_setup.py -f
cd ~

git config --global credential.helper store

chsh -s $(which zsh)

# Docker install
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install -y ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
