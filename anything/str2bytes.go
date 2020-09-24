package anything

import "unsafe"

func str2bytes(s string) []byte {
	tmp := (*[2]uintptr)(unsafe.Pointer(&s))
	str := [3]uintptr{tmp[0], tmp[1], tmp[1]}
	return *(*[]byte)(unsafe.Pointer(&str))
}

func bytes2str(b []byte) string {
	tmp := (*[3]uintptr)(unsafe.Pointer(&b))
	bt := [2]uintptr{tmp[0], tmp[1]}
	return *(*string)(unsafe.Pointer(&bt))
}