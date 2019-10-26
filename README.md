# wifiled
Command line access to generic WIFI LED Controllers

## Usage
This version expects to connect to a single LED Controller that is set by environment variables.

On linux run `export wifiled_ip= 192.168.1.123` and `export wifiled_port = 5577` (port will default to 5577 if not set)

Then you can issue the following commands:

`wifiled on` - turns on LEDs to the last RGBWWW setting

`wifiled off` - turns off LEDs

`wifiled dim 255 0 0` - dims LEDs to Red 255/255 Green 0/255 Blue 0/255 (does not set the Warm White or Cool White, yet)

## To Build
Run `go build` in the root of the directory. 

## Testing
Tested with generic WIFI LED Controllers that use the Magic House mobile app to control.

XCSOURCE DC 12-24V WIFI Remote 5 Channels Controller for iOS Android RGB LED Strip LD686

2014 Protocol - Sends tcp packets to port 5577