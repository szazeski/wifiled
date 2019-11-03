#! /bin/bash

# This script can be setup in your crontab to run hourly with:
# crontab -e
# 0 * * * * ~/cron-script-example.sh

# use tzselect to find your desired timezone, this sets timezone only for
# the duration of the command, but does not change system timezone
DATE=`TZ='America/Chicago' date +%H`
echo date is $DATE

case $DATE in

1)
  # 1AM
  wifiled dim 10
  ;;

7)
  wifiled dim 255
  ;;

9)
  wifiled dim 10
  ;;

17)
  # 5PM
  wifiled dim 255
  ;;

20)
  # 8PM
  wifiled dim 50
  ;;

*)
  ;;

esac
