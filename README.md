# wifiled
Command line access to generic WIFI LED Controllers

Tested with generic WIFI LED Controllers that use the Magic House mobile app to control.

-XCSOURCE DC 12-24V WIFI Remote 5 Channels Controller for iOS Android RGB LED Strip LD686

2014 Protocol

Sends tcp packets to port 5577

## Usage

`wifiled on` - turns on LEDs to full power

`wifiled off` - turns off LEDs

`wifiled dim 255 0 0` - dims LEDs to Red 255/255 Green 0/255 Blue 0/255


###Connecting LED controller
connect to `LEDnetCF643A` (last 5 letters will be different)
3a:59:3f

10.10.123.3
V1.0.14