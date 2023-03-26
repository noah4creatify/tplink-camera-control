package cmd

import "github.com/Luocy7/ha-tplink-camera-control/pkg"

type TurnCommand struct {
	On  bool `long:"on" description:"Turn on the camera and move to preset 1."`
	Off bool `long:"off" description:"Turn off the camera."`
}

func (c *TurnCommand) Execute(_ []string) error {
	cfg, err := pkg.GetConfig()
	if err != nil {
		return err
	}
	r, err := pkg.NewRouter(cfg)
	if err != nil {
		return err
	}
	switch {
	case c.On:
		if err = r.TurnOnCamera(); err != nil {
			return err
		}
	case c.Off:
		if err = r.TurnOffCamera(); err != nil {
			return err
		}
	}
	return nil
}
