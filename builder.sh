mkdir -p /tmp/goalias
sudo chmod -R 0777 /tmp/goalias
rsync -a --exclude .git/ --exclude .gitignore --exclude *.un~ --exclude *.swp --exclude *.swo --exclude *.swn --exclude bin/ "$PWD/." /tmp/goalias
odir="$PWD"
cd /tmp/goalias/src
go build
env GOOS=windows GOARCH=amd64 go build
cd "$odir"
cp /tmp/goalias/src/src "$PWD/bin/goalias"
cp /tmp/goalias/src/src.exe "$PWD/bin/goalias.exe"
