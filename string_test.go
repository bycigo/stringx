package stringx

import "testing"

func TestLen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty string", "", 0},
		{"ascii string", "hello", 5},
		{"chinese characters", "你好世界", 4},
		{"mixed ascii and chinese", "hello你好", 7},
		{"emoji", "👍🎉", 2},
		{"single byte", "a", 1},
		{"single rune", "中", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Len(tt.s); got != tt.want {
				t.Errorf("Len(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"ascii string", "hello", "olleh"},
		{"chinese characters", "你好世界", "界世好你"},
		{"mixed ascii and chinese", "hello你好", "好你olleh"},
		{"palindrome", "abcba", "abcba"},
		{"emoji", "👍🎉", "🎉👍"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.s); got != tt.want {
				t.Errorf("Reverse(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestPadLeft(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		length  int
		padding string
		want    string
	}{
		{"pad with single char", "hello", 10, "*", "*****hello"},
		{"pad with multiple chars", "hello", 10, "ab", "ababahello"},
		{"no padding needed", "hello", 5, "*", "hello"},
		{"length less than string", "hello", 3, "*", "hello"},
		{"empty padding", "hello", 10, "", "hello"},
		{"empty string", "", 5, "*", "*****"},
		{"unicode padding", "hi", 5, "你", "你你你hi"},
		{"unicode string", "你好", 5, "*", "***你好"},
		{"mixed unicode", "你好", 6, "ab", "abab你好"},
		{"multi-byte padding partial", "hi", 4, "你好", "你好hi"},
		{"multi-byte padding partial 2", "hi", 5, "你好世", "你好世hi"},
		{"multi-byte padding partial 3", "hi", 6, "你好世", "你好世你hi"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadLeft(tt.s, tt.length, tt.padding); got != tt.want {
				t.Errorf("PadLeft(%q, %v, %q) = %v, want %v", tt.s, tt.length, tt.padding, got, tt.want)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		length  int
		padding string
		want    string
	}{
		{"pad with single char", "hello", 10, "*", "hello*****"},
		{"pad with multiple chars", "hello", 10, "ab", "helloababa"},
		{"no padding needed", "hello", 5, "*", "hello"},
		{"length less than string", "hello", 3, "*", "hello"},
		{"empty padding", "hello", 10, "", "hello"},
		{"empty string", "", 5, "*", "*****"},
		{"unicode padding", "hi", 5, "你", "hi你你你"},
		{"unicode string", "你好", 5, "*", "你好***"},
		{"mixed unicode", "你好", 6, "ab", "你好abab"},
		{"multi-byte padding partial", "hi", 4, "你好", "hi你好"},
		{"multi-byte padding partial 2", "hi", 5, "你好世", "hi你好世"},
		{"multi-byte padding partial 3", "hi", 6, "你好世", "hi你好世你"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRight(tt.s, tt.length, tt.padding); got != tt.want {
				t.Errorf("PadRight(%q, %v, %q) = %v, want %v", tt.s, tt.length, tt.padding, got, tt.want)
			}
		})
	}
}

func TestPadBoth(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		length  int
		padding string
		want    string
	}{
		{"pad with single char even", "hello", 11, "*", "***hello***"},
		{"pad with single char odd", "hello", 10, "*", "**hello***"},
		{"pad with multiple chars", "hello", 11, "ab", "abahelloaba"},
		{"no padding needed", "hello", 5, "*", "hello"},
		{"length less than string", "hello", 3, "*", "hello"},
		{"empty padding", "hello", 10, "", "hello"},
		{"empty string", "", 6, "*", "******"},
		{"unicode padding", "hi", 6, "你", "你你hi你你"},
		{"unicode string", "你好", 6, "*", "**你好**"},
		{"mixed unicode", "你好", 8, "ab", "aba你好aba"},
		{"multi-byte padding partial both", "hi", 6, "你好", "你好hi你好"},
		{"multi-byte padding partial left", "hi", 5, "你好世", "你hi你好"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadBoth(tt.s, tt.length, tt.padding); got != tt.want {
				t.Errorf("PadBoth(%q, %v, %q) = %v, want %v", tt.s, tt.length, tt.padding, got, tt.want)
			}
		})
	}
}
