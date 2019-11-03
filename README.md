# wifiled
Command line access to generic socket based WIFI LED Controllers

## Usage
This version expects to connect to a single LED Controller that is set by environment variables.

On linux run `export wifiled_ip= 192.168.1.123` and `export wifiled_port = 5577` (port will default to 5577 if not set)

Then you can issue the following commands:

`wifiled on` - turns on LEDs to the last RGBWWW setting

`wifiled off` - turns off LEDs

`wifiled dim 50` - dims all channels to 50/255

`wifiled dim 255 0 0` - dims LEDs to Red 255/255 Green 0/255 Blue 0/255

## To Build
Run `go build` in the root of the directory. 

## Testing
Tested with generic WIFI LED Controllers that use the MagicHome mobile app to control.
 - [MagicHome Android App](https://play.google.com/store/apps/details?id=com.zengge.wifi&hl=en_US)
 - [MagicHome iOS App](https://apps.apple.com/us/app/magic-home-pro/id1187808229)

Devices:
 - XCSOURCE DC 12-24V WIFI Remote 5 Channels Controller for iOS Android RGB LED Strip LD686
    - MPN : HF-LPB100-1
 - LEDENET Smart Wifi LED Controlelr 5 Channels Control 4A4CH
 - (more to come)

2014 Protocol - Sends tcp packets to port 5577


## Other Awesome Projects
Tasmota Firmware : https://github.com/arendst/Tasmota/wiki/MagicHome-LED-strip-controller
ESPurna Firmware : https://github.com/xoseperez/espurna/wiki/Hardware-Magic-Home-LED-Controller
Magichome-python : https://github.com/adamkempenich/magichome-python
node-magichome : https://github.com/jangxx/node-magichome