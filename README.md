# notabar

Because bars are for noobs.

## Description

Display the output of several programs in a system notification window. This
can be used as a mnml information tool showing time, battery, volume, wifi, and
whatever else you want via simple configuration files and running the program.

More info can be found in `notabar(1)` or for the config in `notabar(5)`

Screenshot of configuration and example notification.
![1](https://paste.cf/e3cddc120bf546f6dab93ef3791eeb66798341f5.png)

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
