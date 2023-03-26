package main

import (
	"log"
	"os"

	"github.com/Luocy7/ha-tplink-camera-control/pkg/cmd"
	"github.com/jessevdk/go-flags"
)

type FlagOptions struct {
	Set  cmd.SetCommand  `command:"set" description:"Set the user name and password."`
	Turn cmd.TurnCommand `command:"turn" description:"Turn on or off the camera."`
}

func main() {
	var opt FlagOptions
	p := flags.NewParser(&opt, flags.HelpFlag)
	p.SubcommandsOptional = true

	_, err := p.Parse()

	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 1 {
		p.WriteHelp(os.Stdout)
		return
	}
}
