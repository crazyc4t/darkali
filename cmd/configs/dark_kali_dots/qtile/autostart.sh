#!/bin/bash
# Compositor
picom &
# Resolution
xrandr --output Virtual-1 --mode 1920x1080
# For wallpapers you can use:
# nitrogen --restore
# Or:
feh --bg-scale "/usr/share/backgrounds/kali/kali-ascii-16x9.png"
