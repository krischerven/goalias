package files

import (
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

func TestExists(t *testing.T) {
	autopass(t)
	if !Exists("/usr/bin/ls") {
		t.FailNow()
	}
	if Exists("/usr/bin/notafile") {
		t.FailNow()
	}
}

func TestBinExists(t *testing.T) {
	autopass(t)
	if !BinExists("ls") {
		t.FailNow()
	}
	if BinExists("notafile") {
		t.FailNow()
	}
}

// For this to pass, you need to create /usr/local/bin/goalias-test
func TestUsrLocalBinExists(t *testing.T) {
	autopass(t)
	if !UsrLocalBinExists("goalias-test") {
		t.FailNow()
	}
	if UsrLocalBinExists("notafile") {
		t.FailNow()
	}
}

// For this to pass, you need to create /usr/local/bin/goalias-test
func TestBin(t *testing.T) {
	if Bin("goalias-test") != "/usr/local/bin/goalias-test" {
		t.FailNow()
	}
	if Bin("ls") != "/usr/bin/ls" {
		t.FailNow()
	}
}

func TestDir(t *testing.T) {
	if Dir("/this/is/a/file/path") != "/this/is/a/file/" {
		t.FailNow()
	}
	if Dir("test") != "" {
		t.FailNow()
	}
	if Dir("a/b") != "a/" {
		t.FailNow()
	}
}

func TestRead(t *testing.T) {
	// for now, this will always pass
	// thats because Read() is currently
	// a wrapper around ioutil.ReadFile()
}
