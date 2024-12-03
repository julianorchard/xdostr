# xdostr

A small utility for wrapping [xdotool](https://github.com/jordansissel/xdotool), 
something I find useful on occasion working in an X11 graphical environment 90% 
of the time.

It effectively serves the purpose of a special paste for when an application
isn't happy with whatever your clipboard is using; e.g. pasting into modern
browser applications or a VM that doesn't have clipboard enabled.

## Install

Requires [`xclip`](https://github.com/astrand/xclip) and `xdotool`.

```sh
go install github.com/julianorchard/xdostr@latest
```

Install from source:

```sh
git clone https://github.com/julianorchard/xdostr.git
cd xdostr
make install
```

## Examples

With une pipe:

```sh
echo "Hello, World!" | xdostr
```

With arguments:

```sh
xdostr "Lorum ipsum"
```

From the clipboard!

```sh
xdostr
```

In [i3wm](https://i3wm.org/):

```sh
bindsym --release $mod+Ctrl+v exec "xdostr"
```
