# Crazyc4t
from settings.widgets import screens, widget_defaults, extension_defaults
from settings.keymaps import keys
from settings.vars import *
from settings.groups import groups
from settings.layouts import layouts, floating_layout
from libqtile import hook
import os
import subprocess


@hook.subscribe.startup_once
def autostart():
    home = os.path.expanduser('~/.config/qtile/autostart.sh')
    subprocess.Popen([home])


dgroups_app_rules = []
follow_mouse_focus = True
bring_front_click = False
cursor_warp = False
auto_fullscreen = True
focus_on_window_activation = "smart"
reconfigure_screens = True
auto_minimize = True
wmname = "Qtile"
