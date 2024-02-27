package byteconv

import "unsafe"

func Bytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func String(slice []byte) string {
	return unsafe.String(unsafe.SliceData(slice), len(slice))
}

func Convert[From, To any](value From) To {
	return *(*To)(unsafe.Pointer(&value))
}
