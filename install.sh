#!/bin/sh

system=$(uname -s | awk '{print tolower($0)}')
hardware=$(uname -m)

if [ $hardware = "x86_64" ]
then
  hardware="amd64"
fi

curl -s https://api.github.com/repos/alileza/mnt/releases/latest \
| grep "browser_download_url.*$system-$hardware" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget -qi -

echo "Execute this script to install the binary\n"
echo "   tar -xzvf ./mnt.$system-$hardware.tar.gz -C /usr/local/bin/mnt"
echo "   echo 'alias git='mnt git' >> ~/.bash_profile\n"
