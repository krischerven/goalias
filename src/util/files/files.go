package files

import (
	"io/ioutil"
	"os"
	"strings"
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

func Bin(path string) string {
	if Exists("/usr/local/bin/" + path) {
		return "/usr/local/bin/" + path
	} else if Exists("/usr/bin/" + path) {
		return "/usr/bin/" + path
	} else if Exists("/bin/" + path) {
		return "/bin/" + path
	} else {
		return path
	}
}

func Dir(path string) string {
	split := strings.Split(path, "/")
	ret := ""
	for i := 0; i < len(split)-1; i++ {
		ret += split[i]
		ret += "/"
	}
	return ret
}

func Read(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	return b, err
}
