package bytes

import (
	"bytes"
)

func Split(b []byte, bb []byte) [][]byte {
	return bytes.Split(b, bb)
}

func Contains(b []byte, bb []byte) bool {
	return bytes.Contains(b, bb)
}

func HasPrefix(b []byte, bb []byte) bool {
	return bytes.HasPrefix(b, bb)
}

func Join(b [][]byte, bb []byte) []byte {
	return bytes.Join(b, bb)
}

func Equal(b []byte, bb []byte) bool {
	return bytes.Equal(b, bb)
}

func SplitFirstsLast(b []byte, splitter []byte) [2][]byte {
	ret := bytes.Split(b, splitter)
	switch len(ret) {
	case 0:
		panic("bytes.SplitFirstsLast(): len == 0")
	case 1:
		return [2][]byte{ret[0], nil}
	case 2:
		return [2][]byte{ret[0], ret[1]}
	default:
		arr := [2][]byte{nil, ret[len(ret)-1]}
		for i := 0; i < len(ret)-1; i++ {
			arr[0] = append(arr[0], ret[i]...)
			arr[0] = append(arr[0], splitter...)
		}
		arr[0] = arr[0][0 : len(arr[0])-len(splitter)]
		return arr
	}
}
