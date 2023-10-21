#!/bin/bash

GREEN_COLOR='\033[0;32m'
echo -e"${GREEN_COLOR}Welcome to Installasi For URANUS App"

setRepository(){
    ver_deb=$(cat  /etc/debian_version)
    
    if [[ $ver_deb > "11" ]]; then
         echo "deb http://repo.antix.or.id/debian bullseye main contrib non-free
    deb-src http://repo.antix.or.id/debian bullseye main contrib non-free
    deb http://repo.antix.or.id/debian-security/ bullseye-security main contrib non-free
    deb-src http://repo.antix.or.id/debian-security/ bullseye-security main contrib non-free
    deb http://repo.antix.or.id/debian bullseye-updates main contrib non-free
    deb-src http://repo.antix.or.id/debian bullseye-updates main contrib non-free" >> /etc/apt/sources.list

    else
        echo "deb http://kartolo.sby.datautama.net.id/debian/ buster main contrib non-free
    deb http://kartolo.sby.datautama.net.id/debian/ buster-updates main contrib non-free
    deb http://kartolo.sby.datautama.net.id/debian-security/ buster/updates main contrib non-free" >> /etc/apt/sources.list
   
}

setupDocker() {
    docker-compose up -d 
}

if [ -n "$(uname - a | grep Ubuntu)"]; then
    sudo apt-get update && sudo apt upgrade && sudo apt install docker docker.io docker-compose
elif [ -n "$(uname - a | grep Debian)"];then
    setRepository && apt update -y && apt install docker docker.io docker-compose -y
else
    echo "Just support on debian and Ubuntu"


setupDocker
