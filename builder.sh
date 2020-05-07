mkdir -p /tmp/goalias
sudo chmod -R 0777 /tmp/goalias
sudo cp -r "$PWD/." /tmp/goalias
odir="$PWD"
cd /tmp/goalias/src
go build
env GOOS=windows GOARCH=amd64 go build
cd "$odir"
cp /tmp/goalias/src/src "$PWD/bin/goalias"
cp /tmp/goalias/src/src.exe "$PWD/bin/goalias.exe"
