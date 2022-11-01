#!/bin/bash
echo 'export PATH="/home/$(whoami)/.local/bin:$PATH"' >> ~/.zshrc
cp -r dark_kali_dots/* ~/.config
git clone --depth 1 https://github.com/wbthomason/packer.nvim\
 ~/.local/share/nvim/site/pack/packer/start/packer.nvim