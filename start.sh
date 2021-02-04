#!/bin/bash 
export REDIS_PASSPHRASE="Swift123\$"
cd bin
./Session > ../logs/Session.log 2>&1 &
./Message > ../logs/Message.log 2>&1 &
./File  > ../logs/File.log 2>&1 &
./ERM  > ../logs/ERM.log 2>&1 &
./Certs  > ../logs/Certs.log 2>&1 &
./Bootstrap  > ../logs/Bootstrap.log 2>&1 &
./Data  > ../logs/Data.log 2>&1 &

echo Servers started
