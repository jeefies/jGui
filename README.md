# JGui
**A tiny gui depends on libSDL2**

## Language
I want it to be C at first, but I find it difficult to develop this in C  
So Now it changes into golang, and will keep on developing with golang

## Purpose
To make a simple gui for myself use (to replace tkinter, gtk or qt)  
Notice that it won't be perfect and it would have many issues

## Dependences
It just requires `libSDL2` and `libSDL2-ttf`  
Of course, `libSDL2-gfx libSDL2-mixer libSDL2-net` are also needed in the future but not now.  
> At First, gfx, mixer, net are not used!  
> So you can choose **NOT** to download them


### Debian/Ubuntu
`sudo apt install libsdl2-dev libsdl2-ttf-dev libsdl2-gfx-dev libsdl2-mixer-dev libsdl2-net-dev -y`  
This can automatically download the stable development libs that can use when developing a project.  
###  ArchLinux/Manjaro
`sudo pacman -S sdl2-devel sdl2_{gfx,ttf,mixer,net,image}-devel`  
This can help you download libsdl in your computer too.  
### Centos/Fedora
`sudo yum install SDL2-devel SDL2_{image,ttf,net,mixer,gfx}-devel`  
Also can help you install.
### Windows
See the website of libsdl to find the way to build the dependences.  

## Notice
***We don't use go mod in this package (that means need to turn off GO111MODULE)***  
> type `go env -w GO111MODULE='off'` to turn it off  
> or set `GOPATH` first and `go env -w GO111MODULE='auto'` to let go automatically choose whether use GOPATH or GO MODULE.  
**Maybe when the project is fully developed, i'll change it into newest go mod**
