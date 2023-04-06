# TP-Link IPC44AW Camera Control In GO

## Installation

get release package from [release page](https://github.com/Luocy7/tplink-camera-control/releases)
or build from source

```bash
git clone https://github.com/Luocy7/tplink-camera-control; cd tplink-camera-control
go build -o tpcc
```

## Usage

```bash
$ tpcc -h

Usage:
  tpcc [OPTIONS] [set | turn]

Help Options:
  -h, --help  Show this help message

Available commands:
  set   Set the user name and password.
  turn  Turn on or off the camera.
```

Note that tpcc will create config file `tpcc.toml` in your User Config Directory with `XDG_CONFIG_HOME` or `HOME`
environment variable. detail see [go xdg](https://github.com/adrg/xdg#xdg-base-directory)

## Example

```bash
# set username, password and ip address of camera
tpcc set -a 192.168.9.214 -u admin -p my_passwd

# turn on camera and move to preset 1
tpcc turn --on

# turn off camera
tpcc turn --off
```

## Credits

- [Mecury IPC Control](https://github.com/likaci/mercury-ipc-control)
- [blog.xiazhiri.com](http://blog.xiazhiri.com/Mercury-MIPC251C-4-Reverse.html)
- [ffuf](https://github.com/ffuf/ffuf)