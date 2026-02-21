package stringx

import (
	"unsafe"
)

// FromBytes converts a byte slice to a string without copying the data.
// This is more efficient than using string(bytes) because it avoids unnecessary memory allocation. However, be
// cautious when using this function, as any modification to the bytes will also modify the resulting string, since
// they share the same underlying data. Use this function in appropriate scenarios where you can guarantee that the
// byte slice will not be modified after conversion.
func FromBytes(bytes []byte) string {
	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}

// ToBytes converts a string to a byte slice without copying the data.
// This is more efficient than using []byte(s) because it avoids unnecessary memory allocation. However, be cautious
// when using this function, as any modification to the resulting byte slice will cause a panic, since strings are
// immutable in Go. Use this function in appropriate scenarios where you can guarantee that the resulting byte slice
// will not be modified after conversion.
func ToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
