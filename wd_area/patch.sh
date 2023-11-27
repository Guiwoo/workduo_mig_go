#!/bin/sh

pname=wd-area
#rm -f $pname
#env GOOS=linux GOARCH=amd64 go build -o $pname
#sleep 1
#gtar -cvf $pname.tar $pname
#sleep 1

gtar -cvf area.geojson.tar ./data/area.geojson
sleep 1
scp -i ~/Downloads/oracle.key ./area.geojson.tar opc@158.179.175.176:~/etc/

#scp -i ~/Downloads/oracle.key ./$pname.tar opc@158.179.175.176:~/patch/
#rm -f $pname
#rmf -f $pname.tar