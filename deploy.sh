#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/newweb/
git pull https://github.com/hifuny/devops.git
cd webserver/
./webserver &