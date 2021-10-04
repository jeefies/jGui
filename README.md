# JGui
**A tiny gui depends on libSDL2**

## Language
I want it to be C at first, but I find it difficult to develop this in C  
So Now it changes into golang, and will keep on developing with golang

## Purpose
To make a simple gui for myself use (to replace tkinter, gtk or qt)  
Notice that it won't be perfect and it would have many issues

## Dependences
It just requires `libSDL2`  
Of course, `libSDL2-gfx libSDL2-mixer libSDL2-ttf` are also needed.  
> At First, gfx, mixer are not used!


### Debian/Ubuntu
`sudo apt install libsdl2-dev libsdl2-ttf-dev libsdl2-gfx-dev libsdl2-mixer-dev libsdl2-net-dev -y`  
This can automatically download the stable development libs that can use when developing a project.  
### Windows
See the website of libsdl to find the way to build the dependences.  

## Notice
***We don't use go mod in this package (that means need to turn off GO111MODULE)***  
> type `go env -w GO111MODULE='off'` to turn it off  
**Maybe when the project is fully developed, i'll change it into newest go mod**
