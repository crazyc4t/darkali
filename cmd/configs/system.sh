#!/bin/bash
apt update -yy && apt upgrade -yy
sleep 2
apt install kali-desktop-core git neovim python3 python3-pip libpangocairo-1.0-0 alsa-utils pavucontrol thunar firefox-esr suckless-tools lightdm kitty nitrogen picom golang-go neofetch rofi xorg xserver-xorg xinit python3-pip python3-xcffib python3-cairocffi -yy
sleep 2
pip install qtile
sleep 2
cp qtile.desktop /usr/share/xsessions
sleep 2
systemctl enable lightdm
sleep 2