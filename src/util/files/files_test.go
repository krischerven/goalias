package files

import (
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
	if Bin("goalias-test") != "/usr/local/bin/goalias-test" {
		t.FailNow()
	} else if Bin("ls") != "/usr/bin/ls" {
		t.FailNow()
	}
}

func TestDir(t *testing.T) {
	if Dir("/this/is/a/file/path") != "/this/is/a/file/" {
		t.FailNow()
	} else if Dir("test") != "" {
		t.FailNow()
	} else if Dir("a/b") != "a/" {
		t.FailNow()
	}
}

// For this to pass, you need to create /usr/local/bin/goalias-test
// and make sure it contains the exact text 'hello world\n'
func TestRead(t *testing.T) {
	if b, err := Read("/usr/local/bin/goalias-test"); !bytes.Equal(b, []byte("hello world\n")) ||
		err != nil {
		t.FailNow()
	}
}
