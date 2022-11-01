from libqtile.config import Screen
from libqtile import widget, bar, qtile
from .vars import *

# Widget vars

widget_defaults = dict(font=font, fontsize=18, padding=2, background=theme[0])

extension_defaults = widget_defaults.copy()

# Widgets

screens = [
    Screen(
        top=bar.Bar(
            [
                widget.Sep(
                    linewidth=0,
                    padding=6,
                    foreground=theme[12],
                    background=theme[0],
                ),
                widget.GroupBox(
                    font=font,
                    fontsize=14,
                    margin_y=3,
                    margin_x=0,
                    padding_y=8,
                    padding_x=5,
                    borderwidth=1,
                    highlight_method="text",
                    active=theme[13],
                    inactive=theme[4],
                    rounded=False,
                    urgent_border=theme[0],
                    this_current_screen_border=theme[11],
                    this_screen_border=theme[0],
                    other_current_screen_border=theme[0],
                    other_screen_border=theme[0],
                    disable_drag=True,
                ),
                widget.Sep(
                    linewidth=0,
                    padding=6,
                    foreground=theme[12],
                    background=theme[0],
                ),
                widget.WindowName(
                    fontsize=14, foreground=theme[13], padding=0),
                widget.Systray(background=theme[0], padding=5),
                widget.KeyboardLayout(
                    configured_keyboards=["us", "es olpc"],
                    display_map={"us": "us", "es olpc": "es"},
                    fmt="  {} ",
                    foreground=theme[9],
                    background=theme[0],
                ),
                widget.Volume(
                    foreground=theme[10],
                    background=theme[0],
                    padding=5,
                    mouse_callbacks={
                        "Button1": lambda: qtile.cmd_spawn("pavucontrol")},
                    fmt=" 墳 {} ",
                ),
                widget.TextBox(
                    foreground=theme[11],
                    background=theme[0],
                    fmt=" {} ",
                ),
                widget.Net(
                    interface="eth0",
                    format="  {down} ",
                    background=theme[0],
                    foreground=theme[12],
                    padding=5,
                    mouse_callbacks={
                        "Button1": lambda: qtile.cmd_spawn(terminal + " -e speedtest")
                    },
                ),
                widget.Memory(
                    foreground=theme[13],
                    background=theme[0],
                    mouse_callbacks={
                        "Button1": lambda: qtile.cmd_spawn(terminal + " -e htop")
                    },
                    fmt=" {} ",
                    padding=5,
                ),
                widget.Clock(
                    padding=5,
                    format="  %H:%M ",
                    background=theme[0],
                    foreground=theme[14],
                ),
                widget.CurrentLayoutIcon(
                    background=theme[0], foreground=theme[15]),
            ],
            24,
            opacity=1,
        ),
    ),
]
#
