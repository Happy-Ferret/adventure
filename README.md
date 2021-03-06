# adventure

[![Build Status](https://travis-ci.org/peterhellberg/adventure.svg?branch=master)](https://travis-ci.org/peterhellberg/adventure)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/adventure/game)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/adventure#license-mit)

A small text based adventure game using the [ishell](https://github.com/abiosoft/ishell) package.

I started working on the long flight back to Europe after [GopherCon](http://gophercon.com) 2015.

## Installation

    go get -u github.com/peterhellberg/adventure

## Playing the game

```bash
 ▄▄▄      ▓█████▄  ██▒   █▓▓█████ ███▄    █ ▄▄▄█████▓ █    ██  ██▀███  ▓█████
▒████▄    ▒██▀ ██▌▓██░   █▒▓█   ▀ ██ ▀█   █ ▓  ██▒ ▓▒ ██  ▓██▒▓██ ▒ ██▒▓█   ▀
▒██  ▀█▄  ░██   █▌ ▓██  █▒░▒███  ▓██  ▀█ ██▒▒ ▓██░ ▒░▓██  ▒██░▓██ ░▄█ ▒▒███
░██▄▄▄▄██ ░▓█▄   ▌  ▒██ █░░▒▓█  ▄▓██▒  ▐▌██▒░ ▓██▓ ░ ▓▓█  ░██░▒██▀▀█▄  ▒▓█  ▄
 ▓█   ▓██▒░▒████▓    ▒▀█░  ░▒████▒██░   ▓██░  ▒██▒ ░ ▒▒█████▓ ░██▓ ▒██▒░▒████▒
 ▒▒   ▓▒█░ ▒▒▓  ▒    ░ ▐░  ░░ ▒░ ░ ▒░   ▒ ▒   ▒ ░░   ░▒▓▒ ▒ ▒ ░ ▒▓ ░▒▓░░░ ▒░ ░
  ▒   ▒▒ ░ ░ ▒  ▒    ░ ░░   ░ ░  ░ ░░   ░ ▒░    ░    ░░▒░ ░ ░   ░▒ ░ ▒░ ░ ░  ░
  ░   ▒    ░ ░  ░      ░░     ░     ░   ░ ░   ░       ░░░ ░ ░   ░░   ░    ░
      ░  ░   ░          ░     ░  ░        ░             ░        ░        ░  ░
           ░           ░

Adventure game!
▶ help
You can perform the following commands:
drop, exit, help, items, look, take, use, walk
▶ look
You are standing in the kitchen for the first time.
Paths: garden, livingroom
▶ walk garden
You walked to the garden
```

## License (MIT)

Copyright (c) 2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
