#!/bin/bash
if [ -f /usr/local/bin/goalias ]; then
	sudo rm /usr/local/bin/goalias
	echo "Successfully uninstalled goalias."
else
	echo "Error: goalias is not installed."
fi
