#!/bin/bash
apt update -yy && apt upgrade -yy
sleep 2
apt install kali-desktop-core git neovim python3 python3-pip libpangocairo-1.0-0 alsa-utils pavucontrol thunar firefox-esr suckless-tools lightdm kitty nitrogen picom golang-go neofetch rofi xorg xserver-xorg xinit -yy
sleep 2
systemctl enable --now lightdm
echo -e `export PATH="/home/$(whoami)/.local/bin:$PATH"` >> ~/.zshrc
git clone https://github.com/qtile/qtile.git
cd qtile
pip install .
cd ..
cp qtile.desktop /usr/share/xsessions
cp -r dark_kali_dots/* ~/.config
git clone --depth 1 https://github.com/wbthomason/packer.nvim\
 ~/.local/share/nvim/site/pack/packer/start/packer.nvim