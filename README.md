GoCambridge
===========

GoCambridge is a simple application which enables controlling a Cambridge 340R(and others)
amplifier over the internet/LAN/mobile by sending commands to the amplifiers RS232-port.

Screenshot can be viewed here: http://maep.dk/img/gocambridge.screenshot.png

GoCambridge is written in Googles Go language http://golang.org/

Author: Anders Kusk

Homepage: http://maep.dk/?id=895a9a82295b657049c5118d221c6603

Github: http://github.com/kusk/gocambridge

GoCambridge uses the Go package goserial found here https://github.com/tarm/goserial

To compile and run GoCambridge just:
- Install Go1.1.2 or newer(remember to set GOPATH)
- Install goserial with: **go get github.com/tarm/goserial**
- Download GoCambridge via git or by downloading the zip-file
- Change the TTY in the gocambridge.go source file
- Either run: **go run gocambridge.go** or compile: **go build gocambridge.go**
- Visit 127.0.0.1:8081 and enjoy!
