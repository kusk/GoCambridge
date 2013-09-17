/*
GoCambridge is a simple application which enables controlling a Cambridge 340R(and others)
amplifier over the internet/LAN by sending commands to the amplifiers RS232-port.

GoCambridge is written in Googles Go language.

Author: Anders Kusk
Homeage: http://maep.dk
Github: http://github.com/kusk/gocambridge

GoCambridge uses the Go package goserial found here https://github.com/tarm/goserial

*/

package main

import (
	"fmt"
	"github.com/tarm/goserial"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Packets
	commands := map[string]string{
		"on":         "#1,01,1\r",
		"off":        "#1,01,0\r",
		"dvd":        "#2,01,01\r",
		"video1":     "#2,01,02\r",
		"tuner":      "#2,01,03\r",
		"video2":     "#2,01,04\r",
		"video3":     "#2,01,05\r",
		"tape":       "#2,01,06\r",
		"aux":        "#2,01,07\r",
		"direct":     "#2,01,08\r",
		"analog":     "#2,04,00\r",
		"digital":    "#2,04,01\r",
		"mute":       "#1,11,00\r",
		"volumeup":   "#1,02\r",
		"volumedown": "#1,03\r",
	}

	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600} // Set rs232 port.

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err) // Print hvis ingen tty blev fundet
	}

	bs, err := ioutil.ReadFile("static.html")
	if err != nil {
		log.Fatal(err) // Print hvis static.html mangler
	}

	str := string(bs) // Convert static.html to string

	i := r.URL.Path[1:]

	if i == "" {
		fmt.Fprint(w, str) // Print static.html
	} else {
		for key, value := range commands {
			if key == i {
				s.Write([]byte(value)) // Send packet
				fmt.Fprint(w, str)     // Print static.html
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
