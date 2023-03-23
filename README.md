Tp-link and Mercury cameras Control with GO

## Installation

get release package from [release page](https://github.com/Luocy7/ha-tplink-camera-control/release)
or build from source

```bash
go install github.com/Luocy7/ha-tplink-camera-control 
```

## Usage

```bash
xxx -h
```

## Example

```bash
// add PTZ preset position
{"method":"do","preset":{"set_preset":{"name":"name","save_ptz":"1"}}}

// PTZ to preset position
{"method":"do","preset":{"goto_preset": {"id": "1"}}}

// PTZ by coord
{"method":"do","motor":{"move":{"x_coord":"10","y_coord":"0"}}}

// PTZ horizontal by step
{"method":"do","motor":{"movestep":{"direction":"0"}}}

// PTZ vertical by step
{"method":"do","motor":{"movestep":{"direction":"90"}}}

// stop PTZ
{"method":"do","motor":{"stop":"null"}}

//reset PTZ
{"method":"do","motor":{"manual_cali":"null"}}

// lens mask
{"method":"set","lens_mask":{"lens_mask_info":{"enabled":"on"}}}

// manual alarm
{"method":"do","msg_alarm":{"manual_msg_alarm":{"action":"start"}}}
{"method":"do","msg_alarm":{"manual_msg_alarm":{"action":"stop"}}}

// toggle green led
{"method":"set","led":{"config":{"enabled":"off"}}}
{"method":"set","led":{"config":{"enabled":"on"}}}

//auto track moving obj
{"method":"set","target_track":{"target_track_info":{"enabled":"on"}}}
{"method":"set","target_track":{"target_track_info":{"enabled":"off"}}}

//alarm if found moving obj
{"method":"set","msg_alarm":{"chn1_msg_alarm_info":{"enabled":"on","alarm_type":"0","alarm_mode":["sound"]}}}
{"method":"set","msg_alarm":{"chn1_msg_alarm_info":{"enabled":"on","alarm_type":"0","alarm_mode":["light"]}}}
{"method":"set","msg_alarm":{"chn1_msg_alarm_info":{"enabled":"on","alarm_type":"0","alarm_mode":["sound","light"]}}}
{"method":"set","msg_alarm_plan":{"chn1_msg_alarm_plan":{"enabled":"on","alarm_plan_1":"0000-0000%2c127"}}}

//motion detection
{"method":"set","motion_detection":{"motion_det":{"enabled":"off"}}}
{"method":"set","motion_detection":{"motion_det":{"enabled":"on"}}}
{"method":"set","motion_detection":{"motion_det":{"digital_sensitivity":"20"}}}
{"method":"set","motion_detection":{"motion_det":{"digital_sensitivity":"50"}}}
{"method":"set","motion_detection":{"motion_det":{"digital_sensitivity":"80"}}}

//enable record and plan
{"method":"set","record_plan":{"chn1_channel":{"enabled":"off","monday":"%5b%220000-2400%3a2%22%5d","tuesday":"%5b%220000-2400%3a2%22%5d","wednesday":"%5b%220000-2400%3a2%22%5d","thursday":"%5b%220000-2400%3a2%22%5d","friday":"%5b%220000-2400%3a2%22%5d","saturday":"%5b%220000-2400%3a2%22%5d","sunday":"%5b%220000-2400%3a2%22%5d"}}}
{"method":"set","record_plan":{"chn1_channel":{"enabled":"on","monday":"%5b%220000-2400%3a2%22%5d","tuesday":"%5b%220000-2400%3a2%22%5d","wednesday":"%5b%220000-2400%3a2%22%5d","thursday":"%5b%220000-2400%3a2%22%5d","friday":"%5b%220000-2400%3a2%22%5d","saturday":"%5b%220000-2400%3a2%22%5d","sunday":"%5b%220000-2400%3a2%22%5d"}}}

//reboot and timing reboot
{"method":"do","system":{"reboot":"null"}}
{"method":"set","timing_reboot":{"reboot":{"enabled":"off","day":"7","time":"03%3a00%3a00"}}}
{"method":"set","timing_reboot":{"reboot":{"enabled":"on","day":"7","time":"03%3a00%3a00"}}}

```

Credits to [Mecury IPC Control](https://github.com/likaci/mercury-ipc-control)

ref: http://blog.xiazhiri.com/Mercury-MIPC251C-4-Reverse.html