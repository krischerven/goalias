# goalias
A more flexible alternative to ~/.bash_aliases

### Installation
*goalias does not provide binary releases and it does not support Windows.*

To install goalias, simply run:

./install.sh

### Usage
```bash
sudo goalias set cddocs cd ~/Documents
. cddocs # must be sourced as it changes directory
sudo goalias reset cddocs cd ~/Downloads
. cddocs # we go to ~/Downloads now
```

```bash
sudo goalias set hello echo "Hello, world!"
hello
# Hello, world!
sudo goalias unset hello
hello
# bash: /usr/local/bin/hello: No such file or directory
```
