package pkg

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

const (
	AppName = "tpcc"

	EncryptShortToken = "RDpbLfCPsJZ7fiv"
	EncryptLongToken  = "yLwVl0zKqws7LgKPRQ84Mdt708T1qQ3Ha7xv3H7NyU84p21BriUWBU43odz3iP4rBL3cD02KZciXTysVXiV8ngg6vL48rPJyAUw0HurW20xqxv9aYb4M9wK1Ae0wlro510qXeU07kV57fQMc8L6aLgMLwygtc0F10a0Dg70TOoouyFhdysuRMO51yY5ZlOZZLEal1h0t9YQW0Ko7oBwmCAHoic4HYbUyVeU3sfQ1xtXcPcf1aT303wAQhv66qzW"
)

var (
	CONFIGDIR  = filepath.Join(xdg.ConfigHome, AppName)
	CONFIGFILE = filepath.Join(CONFIGDIR, AppName+".toml")
)
