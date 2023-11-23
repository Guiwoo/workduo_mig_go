#!/bin/sh

pname=wd-area
rm -f $pname
env GOOS=linux GOARCH=amd64 go build -o $pname
sleep 1
gtar -cvf $panme.tar $pname
sleep 1

scp -i ~/Downloads/oracle.key ./$pname opc@158.179.175.176:~/patch/
rm -f $pname