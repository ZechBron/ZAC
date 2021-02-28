package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// set the shit
var m = [4][11]string{
	//     0    1    2    3    4    5    6    7    8    9
	{"Z", " ", ".", ",", "?", "0", "1", "2", "3", "4", "5"},
	{"E", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"},
	{"C", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"},
	{"H", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Nothing to do: usage ./program decode/encode string")
		os.Exit(1)
	}

	fmt.Println(os.Args[1] + " -> " + os.Args[2])

	var out string
	switch os.Args[1] {
	case "decode":
		out = decode(os.Args[2])
	case "encode":
		out = encode(os.Args[2])
	}
	fmt.Println(out)
}

func decode(decode string) string {
	c := splitme(strings.ToUpper(decode), 2)
	var p string
	for i := 0; i < len(c); i++ {
		for a := 0; a < 4; a++ {
			for s := 0; s < 11; s++ {
				if c[i] == m[a][0]+strconv.Itoa(s) {
					p += m[a][s+1]
				}
			}
		}
	}
	return p
}

func encode(encode string) string {
	c := strings.Split(strings.ToUpper(encode), "")
	var p string
	for i := 0; i < len(c); i++ {
		for a := 0; a < 4; a++ {
			for s := 1; s < 11; s++ {
				if c[i] == m[a][s] {
					o := getletter(strconv.Itoa(a))
					p += o + strconv.Itoa(s-1)
				}
			}
		}
	}
	return p
}

// too lazy to make better
func getletter(c string) string {
	var p string
	switch c {
	case "0":
		p = "Z"
	case "1":
		p = "E"
	case "2":
		p = "C"
	case "3":
		p = "H"
	}
	return p
}

func splitme(f string, n int) []string {
	a := ""
	s := []string{}
	c := bytes.Runes([]byte(f))
	l := len(c)

	for i, r := range c {
		a = a + string(r)
		if (i+1)%n == 0 {
			s = append(s, a)
			a = ""
		} else if (i + 1) == l {
			s = append(s, a)
		}
	}
	return s
}
