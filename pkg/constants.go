package pkg

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

const (
	AppName           = "tpcc"
	DefaultUserName   = "admin"
	EncryptShortToken = "RDpbLfCPsJZ7fiv"
	EncryptLongToken  = "yLwVl0zKqws7LgKPRQ84Mdt708T1qQ3Ha7xv3H7NyU84p21BriUWBU43odz3iP4rBL3cD02KZciXTysVXiV8ngg6vL48rPJyAUw0HurW20xqxv9aYb4M9wK1Ae0wlro510qXeU07kV57fQMc8L6aLgMLwygtc0F10a0Dg70TOoouyFhdysuRMO51yY5ZlOZZLEal1h0t9YQW0Ko7oBwmCAHoic4HYbUyVeU3sfQ1xtXcPcf1aT303wAQhv66qzW"
	ipc44awN          = "C861E9DE714597CC9F62DC68E0722306BE9CD45DA9AE1C6B9E2C83B918466D380904A0108B2830B834DD9A63A6BE4F2D6F40FBC93F5AF583D7E37346CA9EEF8D"
	ipc44awE          = "10001"
)

var (
	CONFIGDIR  = filepath.Join(xdg.ConfigHome, AppName)
	CONFIGFILE = filepath.Join(CONFIGDIR, AppName+".toml")
)

const (
	PostLoginUrl           = "http://%s/"
	PostDsUrl              = "http://%s/stok=%s/ds"
	PayloadLogin           = `{"method":"do","login":{"username":"%s","encrypt_type":"2","password":"%s","md5_encrypt_type": "1"}}`
	PayloadGetBasicInfo    = `{"method":"get","device_info":{"name":["basic_info"]}}`                       // ok
	PayloadGetModuleSpec   = `{"method":"get","function":{"name":["module_spec"]}}`                         // ok
	PayloadGetHarddisk     = `{"method":"get","harddisk_manage":{"table":["hd_info"],"name":["harddisk"]}}` // ok
	PayloadGetNetworkType  = `{"method":"do","network":{"get_connection_type":"null"}}`                     // ok
	PayloadGetClockStatus  = `{"method":"get","system":{"name":["clock_status"]}}`                          // ok
	PayloadGetMediaEncrypt = `{"method":"get","cet":{"name":["media_encrypt"]}}`                            // ok
	PayloadGetVhttpd       = `{"method":"get","cet":{"name":["vhttpd"]}}`                                   // ok
	PayloadGetAll          = `{
		"method":"get",
		"harddisk_manage":{"table":["hd_info"],"name":["harddisk"]},
		"wlan":{"name":["default_ap"]},
		"msg_alarm":{"name":["chn1_msg_alarm_info"]},
		"motion_detection":{"name":["motion_det"]},
		"tamper_detection":{"name":["tamper_det"]},
		"record_plan":{"name":["chn1_channel"]},
		"greeter":{"name":["chn1_greeter_ctrl"]},
		"cet":{"name":["media_encrypt"]}
	}`
	PayloadGetAll2 = `{
		"method":"get",
		"led":{"name":["config"]},
		"target_track":{"name":["target_track_info"]},
		"audio_config":{"name":["microphone","speaker"]},
		"image":{"name":["switch","common"]}
	}`
	PayloadGetHomeassistant = `{
		"method":"get",
		"led":{"name":["config"]},
		"target_track":{"name":["target_track_info"]},
		"lens_mask":{"name":["lens_mask_info"]},
		"preset":{"name":["preset"]},
		"msg_alarm":{"name":["chn1_msg_alarm_info"]}
	}`

	PayloadCheckUpgrade = `{"method":"do","cloud_config":{"check_fw_version":"null"}}` // ok

	PayloadGetLed    = `{"method":"get","led":{"name":["config"]}}`          // ok
	PayloadSetLedOn  = `{"method":"set","led":{"config":{"enabled":"on"}}}`  // ok
	PayloadSetLedOff = `{"method":"set","led":{"config":{"enabled":"off"}}}` // ok

	PayloadGetTrack    = `{"method":"get","target_track":{"name":["target_track_info"]}}`          // ok
	PayloadSetTrackOn  = `{"method":"set","target_track":{"target_track_info":{"enabled":"on"}}}`  // ok
	PayloadSetTrackOff = `{"method":"set","target_track":{"target_track_info":{"enabled":"off"}}}` // ok

	PayloadGetVideoFlip    = `{"method":"get","image":{"name":["switch"]}}`
	PayloadSetVideoFlipOn  = `{"method":"set","image":{"switch":{"flip_type":"center"}}}`
	PayloadSetVideoFlipOff = `{"method":"set","image":{"switch":{"flip_type":"off"}}}`

	PayloadGetAlarmPlan = `{"method":"get","msg_alarm_plan":{"name":["chn1_msg_alarm_plan"]}}`
	PayloadSetAlarmOn   = `{"method":"set","msg_alarm":{"chn1_msg_alarm_info":{"enabled":"on"}}}`
	PayloadSetAlarmOff  = `{"method":"set","msg_alarm":{"chn1_msg_alarm_info":{"enabled":"off"}}}`
	PayloadAlarmStart   = `{"method":"do","msg_alarm":{"manual_msg_alarm":{"action":"start"}}}`
	PayloadAlarmStop    = `{"method":"do","msg_alarm":{"manual_msg_alarm":{"action":"stop"}}}`

	PayloadGetLensmask    = `{"method":"get","lens_mask":{"name":["lens_mask_info"]}}`
	PayloadSetLensmaskOn  = `{"method":"set","lens_mask":{"lens_mask_info":{"enabled":"on"}}}`
	PayloadSetLensmaskOff = `{"method":"set","lens_mask":{"lens_mask_info":{"enabled":"off"}}}`

	PayloadGetPreset    = `{"method":"get","preset":{"name":["preset"]}}`
	PayloadSetPreset    = `{"method":"do","preset":{"set_preset":{"name":"%s","save_ptz":"1"}}}`
	PayloadGotoPreset   = `{"method":"do","preset":{"goto_preset": {"id": "%s"}}}`
	PayloadDeletePreset = `{"method":"do","preset":{"remove_preset":{"id":["%s"]}}}`

	PayloadDoMotorStop      = `{"method":"do","motor":{"stop":"null"}}`
	PayloadDoMotorStepLeft  = `{"method":"do","motor":{"movestep":{"direction":"180"}}}`
	PayloadDoMotorStepRight = `{"method":"do","motor":{"movestep":{"direction":"0"}}}`
	PayloadDoMotorStepUp    = `{"method":"do","motor":{"movestep":{"direction":"90"}}}`
	PayloadDoMotorStepDown  = `{"method":"do","motor":{"movestep":{"direction":"270"}}}`
	PayloadDoMotorMove      = `{"method":"do","motor":{"move":{"x_coord":"%s","y_coord":"%s"}}}`
	PayloadDoMotorMoveLeft  = `{"method":"do","motor":{"move":{"x_coord":"-10","y_coord":"0"}}}`
	PayloadDoMotorMoveRight = `{"method":"do","motor":{"move":{"x_coord":"10","y_coord":"0"}}}`
	PayloadDoMotorMoveUp    = `{"method":"do","motor":{"move":{"x_coord":"0","y_coord":"10"}}}`
	PayloadDoMotorMoveDown  = `{"method":"do","motor":{"move":{"x_coord":"0","y_coord":"-10"}}}`
)
