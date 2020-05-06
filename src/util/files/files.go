package files

import (
	"io/ioutil"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	} else {
		// panic: stat PATH: no such file or directory
		return false
	}
}

func BinExists(path string) bool {
	for _, pfx := range []string{
		"/bin/",
		"/usr/bin/",
		"/usr/local/bin/",
	} {
		if Exists(pfx + path) {
			return true
		}
	}
	return false
}

func UsrLocalBinExists(path string) bool {
	return Exists("/usr/local/bin/" + path)
}

func Read(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	return string(b), err
}