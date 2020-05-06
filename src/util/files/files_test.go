package files

import (
	"testing"
)

//
// Note: these tests will all fail if you aren't on a Unix system.
//

func TestExists(t *testing.T) {
	if !Exists("/usr/bin/ls") {
		t.FailNow()
	}
	if Exists("/usr/bin/notafile") {
		t.FailNow()
	}
}

func TestBinExists(t *testing.T) {
	if !BinExists("ls") {
		t.FailNow()
	}
	if BinExists("notafile") {
		t.FailNow()
	}
}

// For this to pass, you need to create /usr/local/bin/goalias-test
func TestUsrLocalBinExists(t *testing.T) {
	if !UsrLocalBinExists("goalias-test") {
		t.FailNow()
	}
	if UsrLocalBinExists("notafile") {
		t.FailNow()
	}
}
