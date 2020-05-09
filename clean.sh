#!/bin/bash
# This file is designed for debugging:
# only use it if you know what you're doing and you don't mind losing all of your aliases
# if you just want to clear your aliases, use 'goalias unset !'
if [ -f "/etc/goalias/registry.txt" ] && [ $(wc -l < /etc/goalias/registry.txt) != 0 ]; then
	read -p "Registry is not empty; continue? (Y/N)" -n 1 -r
	echo
	if ! [[ $REPLY =~ ^[Yy]$ ]]
	then
		exit 0
	fi
fi

sudo rm -rf /tmp/goalias
sudo rm -rf /etc/goalias
echo "cleaned"
