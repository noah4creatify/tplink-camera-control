package cmd

import (
	"github.com/Luocy7/ha-tplink-camera-control/pkg"
)

type SetCommand struct {
	UserName string `short:"u" long:"user" description:"The name of the user."`
	PassWord string `short:"p" long:"password" description:"The password of the user."`
	Address  string `short:"a" long:"address" description:"The ip address of the camera."`
	Option   *pkg.ConfigOptions
}

func (c *SetCommand) Execute(_ []string) error {
	opt, err := pkg.ReadConfig()
	if err != nil {
		return err
	}
	if c.UserName != "" {
		opt.UserName = c.UserName
	}
	if c.PassWord != "" {
		opt.PassWord = c.PassWord
	}
	if c.Address != "" {
		opt.Address = c.Address
	}
	if err = pkg.WriteConfig(opt); err != nil {
		return err
	}
	c.Option = opt
	return nil
}
