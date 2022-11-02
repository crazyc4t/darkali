#!/bin/bash
apt update -yy && apt upgrade -yy
apt install kali-desktop-core git neovim python3 python3-pip libpangocairo-1.0-0 alsa-utils pavucontrol thunar suckless-tools lightdm kitty nitrogen picom golang-go neofetch rofi xorg xserver-xorg xinit python3-pip python3-xcffib python3-cairocffi keepassxc feh -yy
pip install qtile
cp ./cmd/configs/qtile.desktop /usr/share/xsessions
mv /etc/lightdm/lightdm-gtk-greeter.conf /etc/lightdm/lightdm-config-original.conf
cp ./cmd/configs/lightdm-gtk-greeter.conf /etc/lightdm/
systemctl enable lightdm
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
apt update -yy
apt install librewolf -y
apt remove qterminal -yy