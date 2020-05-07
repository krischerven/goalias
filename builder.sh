#!/bin/bash
mkdir -p /tmp/goalias
sudo chmod -R 0777 /tmp/goalias

# copy important files
rsync \
	-a \
	--exclude .git/ --exclude .gitignore --exclude *.un~ --exclude *.swp \
	--exclude *.swo --exclude *.swn --exclude bin/ "$PWD/." /tmp/goalias

odir="$PWD"
cd /tmp/goalias/src
if [ -f src ]; then
	rm src
fi
if [ -f src.exe ]; then
	rm src.exe
fi
go build
env GOOS=windows GOARCH=amd64 go build
cd "$odir"
if [ -f /tmp/goalias/src/src ]; then
	cp /tmp/goalias/src/src "$PWD/bin/goalias"
fi
if [ -f /tmp/goalias/src/src.exe ]; then
	cp /tmp/goalias/src/src.exe "$PWD/bin/goalias.exe"
fi
