#!/bin/bash
apt update -yy && apt upgrade -yy
sleep 2
apt install kali-desktop-core git neovim python3 python3-pip libpangocairo-1.0-0 alsa-utils pavucontrol thunar suckless-tools lightdm kitty nitrogen picom golang-go neofetch rofi xorg xserver-xorg xinit python3-pip python3-xcffib python3-cairocffi keepassxc feh -yy
sleep 2
pip install qtile
sleep 2
cp ./cmd/configs/qtile.desktop /usr/share/xsessions
sleep 2
mv /etc/lightdm/lightdm-gtk-greeter.conf /etc/lightdm/lightdm-config-original.conf
sleep 1
cp ./cmd/configs/lightdm-gtk-greeter.conf /etc/lightdm/
sleep 2
systemctl enable lightdm
sleep 2
distro=$(if echo " bullseye focal impish jammy uma una vanessa" | grep -q " $(lsb_release -sc) "; then echo $(lsb_release -sc); else echo focal; fi)
wget -O- https://deb.librewolf.net/keyring.gpg | sudo gpg --dearmor -o /usr/share/keyrings/librewolf.gpg
tee /etc/apt/sources.list.d/librewolf.sources << EOF > /dev/null
Types: deb
URIs: https://deb.librewolf.net
Suites: $distro
Components: main
Architectures: amd64
Signed-By: /usr/share/keyrings/librewolf.gpg
EOF
apt update
apt install librewolf -y
apt remove qterminal
sleep 2