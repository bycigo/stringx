package stringx_test

import (
	"strings"
	"testing"

	"github.com/bycigo/stringx"
)

func TestFromBytes(t *testing.T) {
	tests := []struct {
		name string
		arg  []byte
		want string
	}{
		{"convert nil byte slice", nil, ""},
		{"convert empty byte slice", []byte{}, ""},
		{"convert byte slice to string", []byte("hello"), "hello"},
		{"convert unicode byte slice", []byte("你好世界"), "你好世界"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringx.FromBytes(tt.arg); got != tt.want {
				t.Errorf("FromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

var fromBytesResult string

func BenchmarkFromBytes(b *testing.B) {
	// 10MB
	data := make([]byte, 1024*1024*10)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fromBytesResult = stringx.FromBytes(data)
	}
}

func BenchmarkBytesToString(b *testing.B) {
	// 10MB
	data := make([]byte, 1024*1024*10)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fromBytesResult = string(data)
	}
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []byte
	}{
		{"convert empty string", "", []byte{}},
		{"convert string to byte slice", "hello", []byte("hello")},
		{"convert unicode string", "你好世界", []byte("你好世界")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringx.ToBytes(tt.arg); string(got) != string(tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

var toBytesResult []byte

func BenchmarkToBytes(b *testing.B) {
	// 10MB
	s := strings.Repeat("a", 1024*1024*10)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toBytesResult = stringx.ToBytes(s)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	// 10MB
	s := strings.Repeat("a", 1024*1024*10)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toBytesResult = []byte(s)
	}
}
