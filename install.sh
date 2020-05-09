#!/bin/bash
(rm bin/goalias) > /dev/null 2>&1
./builder.sh --silent
if [ -f bin/goalias ]; then
	(sudo rm /usr/local/bin/goalias) > /dev/null 2>&1
	sudo ln -s "$PWD/bin/goalias" /usr/local/bin/goalias
	echo "Successfully installed goalias."
else
	echo "Error installing goalias." | perl -wln -M'Term::ANSIColor' -e 'print "\e[1;91m", "$_", "\e[0m"'
fi
