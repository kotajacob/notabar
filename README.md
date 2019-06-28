# notabar

Because bars are for noobs.

## Description

Display the output of several programs in a system notification window. This
can be used as a mnml information tool showing time, battery, volume, wifi, and
whatever else you want at the press of a buttom.

More info can be found in notabar(1) or for the config in notabar(5)

## Config

The config file for notabar is where you specify which programs provide text to
be displayed by notabar. The config file is a simple csv file which by default
is located in `~/.config/notabar/config` though if you've changed your
XDG_CONFIG_HOME it will follow that standard.

Below is an example config.

```
txt,ti:
cmd,date,+%H:%M
txt,\n

txt,ba:
cmd,battery
cmd,battery,+
txt,\n

txt,li:
cmd,backlight
txt,\n
```

It will render something like this.

```
ti:14:07
ba:76+
li:70
```

There are two kinds of entries. `txt` which just prints the following text and
`cmd` which prints the STDOUT from a provided command. Blank lines are ignored
so I've used them to seperate groups on config based on which line they modify
in the output. Most programs return a newline at the end of their STDOUT, but
notabar strips them to allow for more granular control. You can simply put a
`txt,\n` when you want a newline. Line wrapping is of course controlled by your
notification daemon so investigate its config if weird line wrapping occurs.

In that example config the battery and backlight programs are also small
utilities I've developed. They can be found in my collection of small programs
here https://git.sr.ht/~kota/useless

More info can be found in notabar(1) or for the config in notabar(5)

## Build

Build dependencies  

 * golang
 * make
 * sed
 * scdoc

`make all`

## Install

Optionally configure `config.mk` to specify a different install location.  
Defaults to `/usr/local/`

`sudo make install`

## Uninstall

`sudo make uninstall`
