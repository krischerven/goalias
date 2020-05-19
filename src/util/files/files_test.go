package files

import (
	"fmt"
	"github.com/krischerven/goalias/src/util/bytes"
	"runtime"
	"testing"
)

//
// Note: these tests will all pass if you aren't on a Unix system.
//
func autopass(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.SkipNow()
	}
}

func failstr(t *testing.T, want string, got string) {
	defer t.FailNow()
	fmt.Printf("Test failure (want: %s, got: %s)\n", want, got)
}

func TestExists(t *testing.T) {
	autopass(t)
	if !Exists("/usr/bin/ls") {
		t.FailNow()
	} else if Exists("/usr/bin/notafile") {
		t.FailNow()
	}
}

func TestBinExists(t *testing.T) {
	autopass(t)
	if !BinExists("ls") {
		t.FailNow()
	} else if BinExists("notafile") {
		t.FailNow()
	}
}

// For this to pass, you need to create /usr/local/bin/goalias-test
func TestUsrLocalBinExists(t *testing.T) {
	autopass(t)
	if !UsrLocalBinExists("goalias-test") {
		t.FailNow()
	} else if UsrLocalBinExists("notafile") {
		t.FailNow()
	}
}

// For this to pass, you need to create /usr/local/bin/goalias-test
func TestBin(t *testing.T) {
	autopass(t)
	if Bin("goalias-test") != "/usr/local/bin/goalias-test" {
		failstr(t, "/usr/local/bin/goalias-test", Bin("goalias-test"))
	} else if Bin("ls") != "/usr/bin/ls" {
		failstr(t, "/usr/bin/ls", Bin("ls"))
	}
}

func TestDir(t *testing.T) {
	autopass(t)
	if Dir("/this/is/a/file/path") != "/this/is/a/file/" {
		failstr(t, "/this/is/a/file/", Dir("/this/is/a/file/path"))
	} else if Dir("test") != "" {
		failstr(t, "", Dir("test"))
	} else if Dir("a/b") != "a/" {
		failstr(t, "a/", Dir("a/b"))
	}
}

// For this to pass, you need to create /usr/local/bin/goalias-test
// and make sure it contains the exact text 'hello world\n'
func TestRead(t *testing.T) {
	autopass(t)
	if b, err := Read("/usr/local/bin/goalias-test"); !bytes.Equal(b, []byte("hello world\n")) ||
		err != nil {
		if err == nil {
			failstr(t, string([]byte("hello world\n")), string(b))
		} else {
			t.FailNow()
		}
	}
}
