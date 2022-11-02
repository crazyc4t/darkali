#!/bin/bash
echo 'export PATH="/home/$(whoami)/.local/bin:$PATH"' >> ~/.zshrc
cp -r ./cmd/configs/dark_kali_dots/* ~/.config
pip install psutil
git clone --depth 1 https://github.com/wbthomason/packer.nvim ~/.local/share/nvim/site/pack/packer/start/packer.nvim
./cmd/configs/nerdy.sh Hack