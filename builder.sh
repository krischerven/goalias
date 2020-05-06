mkdir -p /tmp/goalias
sudo cp -r "$PWD/." /tmp/goalias
odir="$PWD"
cd /tmp/goalias
go build
env GOOS=windows GOARCH=amd64 go build
cd "$odir"
cp /tmp/goalias/goalias "$PWD/bin"
cp /tmp/goalias/goalias.exe "$PWD/bin"
