#!/bin/bash
echo 'export PATH="/home/$(whoami)/.local/bin:$PATH"' >> ~/.zshrc
sleep 2
cp -r dark_kali_dots/* ~/.config
sleep 2
pip install psutil
sleep 2
git clone --depth 1 https://github.com/wbthomason/packer.nvim\
~/.local/share/nvim/site/pack/packer/start/packer.nvim
sleep 2
./nerdy.sh Hack
sleep 2