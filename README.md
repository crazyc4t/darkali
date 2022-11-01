# DarkArch Installation Script

- kali-desktop-core git
- sudo apt update && sudo apt upgrade
  sudo apt update -yy

sudo apt install python3 python3-pip python-gobject python-dbus libpangocairo-1.0-0 alsa-utils pavucontrol pcmanfm firefox-esr suckless-tools

pip3 install xcffib

pip3 install --no-cache-dir cairocffi

git clone https://github.com/qtile/qtile.git
cd qtile

pip3 install .

then add

qtile start

to your ~/.xinitrc
