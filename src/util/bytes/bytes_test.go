package bytes

import (
	native "bytes"
	"testing"
)

func TestSplit(t *testing.T) {
	if !Equal(Split([]byte("this is a sentence"), []byte(" "))[2],
		native.Split([]byte("this is a sentence"), []byte(" "))[2]) {
		t.FailNow()
	}
}

func TestContains(t *testing.T) {
	if Contains([]byte("this is a sentence"), []byte("is a ")) !=
		native.Contains([]byte("this is a sentence"), []byte("is a ")) {
		t.FailNow()
	}
}

func TestHasPrefix(t *testing.T) {
	if HasPrefix([]byte("this is a sentence"), []byte("this is ")) !=
		native.HasPrefix([]byte("this is a sentence"), []byte("this is ")) {
		t.FailNow()
	}
}

func TestJoin(t *testing.T) {
	if !Equal(
		Join([][]byte{[]byte("a"), []byte("b"), []byte("c")}, []byte(" ")),
		native.Join([][]byte{[]byte("a"), []byte("b"), []byte("c")}, []byte(" ")),
	) {
		t.FailNow()
	}
}

func TestEqual(t *testing.T) {
	if Equal([]byte("abc"), []byte("abc")) != native.Equal([]byte("abc"), []byte("abc")) {
		t.FailNow()
	} else if Equal([]byte("abc"), []byte("def")) != native.Equal([]byte("abc"), []byte("def")) {
		t.FailNow()
	}
}

func TestSplitFirstsLast(t *testing.T) {
	arr := SplitFirstsLast([]byte("test"), []byte("x"))
	if !Equal(arr[0], []byte("test")) || arr[1] != nil {
		t.FailNow()
	} else if arr = SplitFirstsLast([]byte("axb"), []byte("x")); !Equal(arr[0], []byte("a")) ||
		!Equal(arr[1], []byte("b")) {
		t.FailNow()
	} else if arr = SplitFirstsLast([]byte("axbxc"), []byte("x")); !Equal(arr[0], []byte("axb")) ||
		!Equal(arr[1], []byte("c")) {
		t.FailNow()
	}
}
