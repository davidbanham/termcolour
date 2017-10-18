package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"regexp"
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

	confPath := usr.HomeDir + "/.config/alacritty/alacritty.yml"

	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`colors: \*.*\n`)
	replaced := reg.ReplaceAll(buf, []byte("colors: *"+modes[targetColour]+"\n"))

	err = ioutil.WriteFile(confPath, replaced, 0664)
	if err != nil {
		log.Fatal(err)
	}
}

var modes = map[string]string{
	"dark":  "dark",
	"d":     "dark",
	"light": "light",
	"l":     "light",
}
