#!/bin/bash

# run this command in terminal before chmod +x instruction.sh
sudo apt-get update
#git
sudo apt install git
#vscode
sudo snap install --classic code
#docker 
sudo apt-get update


sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo apt-key fingerprint 0EBFCD88



sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"


sudo apt update

apt-cache policy docker-ce


sudo apt-get install docker-ce docker-ce-cli containerd.io

#Docker verificator
sudo docker run hello-world