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

const (
	PayloadGetKey = `{"method": "do", "login": {}}`
	PayloadLogin  = `{"method":"do","login":{"username":"admin","encrypt_type":"2","password":"%s"}}`

	PayloadGetBasicInfo    = `{"method":"get","device_info":{"name":["basic_info"]}}`
	PayloadGetModuleSpec   = `{"method":"get","function":{"name":["module_spec"]}}`
	PayloadGetCapabilities = `{"method":"get","audio_capability":{"name":["device_speaker","device_microphone"]},"motor":{"name":["capability"]},"playback":{"table":["scale_capability"]},"cet":{"name":["media_encrypt"]}}`
	PayloadGetHarddisk     = `{"method":"get","harddisk_manage":{"table":["hd_info"],"name":["harddisk"]}}`
	PayloadGetNetworkType  = `{"method":"do","network":{"get_connection_type":"null"}}`
	PayloadGetClockStatus  = `{"method":"get","system":{"name":["clock_status"]}}`
	PayloadGetMediaEncrypt = `{"method":"get","cet":{"name":["media_encrypt"]}}`
	PayloadGetVhttpd       = `{"method":"get","cet":{"name":["vhttpd"]}}`
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
		"msg_alarm":{"name":["chn1_msg_alarm_info"]},
	}`

	PayloadCheckUpgrade = `{"method":"do","cloud_config":{"check_fw_version":"null"}}`

	PayloadGetLed    = `{"method":"get","led":{"name":["config"]}}`
	PayloadSetLedOn  = `{"method":"set","led":{"config":{"enabled":"on"}}}`
	PayloadSetLedOff = `{"method":"set","led":{"config":{"enabled":"off"}}}`

	PayloadGetTrack    = `{"method":"get","target_track":{"name":["target_track_info"]}}`
	PayloadSetTrackOn  = `{"method":"set","target_track":{"target_track_info":{"enabled":"on"}}}`
	PayloadSetTrackOff = `{"method":"set","target_track":{"target_track_info":{"enabled":"off"}}}`

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
	PayloadGotoPreset   = `{"method":"do","preset":{"goto_preset": {"id": "1"}}}`
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
