package cmd

import (
	"fmt"

	"github.com/Luocy7/ha-tplink-camera-control/pkg"
)

type SetCommand struct {
	UserName string `short:"u" long:"user" description:"The name of the user."`
	PassWord string `short:"p" long:"password" description:"The password of the user."`
	Option   *pkg.ConfigOptions
}

func (c *SetCommand) Execute(_ []string) error {
	if c.UserName == "" {
		return fmt.Errorf("user name is required")
	}
	if c.PassWord == "" {
		return fmt.Errorf("password is required")
	}

	c.Option.UserName = c.UserName
	c.Option.PassWord = c.PassWord
	return nil
}
