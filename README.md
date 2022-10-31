# DarkArch Installation Script

- kali-desktop-core
- sudo apt update && sudo apt upgrade
- sudo apt install gcc make
sudo apt update -yy

sudo apt install git vim xorg xserver-xorg

sudo apt install python3 python3-pip python-gobject python-dbus libpangocairo-1.0-0 alsa-utils pavucontrol pcmanfm firefox-esr suckless-tools

pip3 install xcffib

pip3 install --no-cache-dir cairocffi

git clone git://github.com/qtile/qtile.git

cd qtile

pip3 install .

then add

qtile start

to your ~/.xinitrc