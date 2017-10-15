package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("What colour, dickhead?")
		os.Exit(1)
	}

	targetColour := os.Args[1]

	if modes[targetColour] == "" {
		fmt.Println("I dunno that one.")
		os.Exit(1)
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	confPath := usr.HomeDir + "/.config/xfce4/terminal/terminalrc"

	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}

	str := string(buf)

	lines := strings.Split(str, "\n")

	goodLines := []string{}

	for _, line := range lines {
		if len(line) > 16 {
			if line[:16] == "ColorBackgroundV" {
				goodLines = append(goodLines, line)
				continue
			}
		}
		if len(line) > 4 {
			if line[:4] == "Name" {
				continue
			}
		}
		if len(line) > 12 {
			switch line[:12] {
			case "ColorPalette":
				continue
			case "ColorForegro":
				continue
			case "ColorBackgro":
				continue
			case "ColorCursor=":
				continue
			default:
				goodLines = append(goodLines, line)
			}
		} else {
			goodLines = append(goodLines, line)
		}
	}

	goodLines = append(goodLines, modes[targetColour])

	joined := strings.Join(goodLines, "\n")

	err = ioutil.WriteFile(confPath, []byte(joined), 0664)
	if err != nil {
		log.Fatal(err)
	}
}

var dark = `
Name=Base16-Spacemacs
ColorPalette=#1f2022;#f2241f;#67b11d;#b1951d;#4f97d7;#a31db1;#2d9574;#a3a3a3;#585858;#f2241f;#67b11d;#b1951d;#4f97d7;#a31db1;#2d9574;#f8f8f8
ColorCursor=x
ColorForeground=#a3a3a3
ColorBackground=#1f2022
`

var light = `
Name=Base16-Summerfruit Light
ColorForeground=#101010
ColorBackground=#FFFFFF
ColorCursor=x
ColorPalette=#FFFFFF;#FF0086;#00C918;#ABA800;#3777E6;#AD00A1;#1FAAAA;#101010;#B0B0B0;#FF0086;#00C918;#ABA800;#3777E6;#AD00A1;#1FAAAA;#202020
`

var modes = map[string]string{
	"dark":  dark,
	"d":     dark,
	"light": light,
	"l":     light,
}
