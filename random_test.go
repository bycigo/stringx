package stringx_test

import (
	"strings"
	"sync"
	"testing"

	"github.com/bycigo/stringx"
)

func TestNewRandomFactory(t *testing.T) {
	t.Run("default options", func(t *testing.T) {
		factory := stringx.NewRandomFactory()
		if factory == nil {
			t.Fatal("expected non-nil factory")
		}
	})

	t.Run("custom alphabet", func(t *testing.T) {
		customAlphabet := "abc123"
		factory := stringx.NewRandomFactory(stringx.RandomAlphabet(customAlphabet))
		if factory == nil {
			t.Fatal("expected non-nil factory")
		}

		// Generate a string and verify it only contains characters from the custom alphabet
		str := factory.MakeRandomString(100)
		for _, c := range str {
			if !strings.ContainsRune(customAlphabet, c) {
				t.Errorf("generated string contains character %c not in alphabet %s", c, customAlphabet)
			}
		}
	})
}

func TestRandomFactory_MakeRandomString(t *testing.T) {
	factory := stringx.NewRandomFactory()

	t.Run("generates string of correct length", func(t *testing.T) {
		lengths := []int{1, 5, 10, 50, 100}
		for _, length := range lengths {
			str := factory.MakeRandomString(length)
			if len(str) != length {
				t.Errorf("expected length %d, got %d", length, len(str))
			}
		}
	})

	t.Run("generates different strings", func(t *testing.T) {
		str1 := factory.MakeRandomString(20)
		str2 := factory.MakeRandomString(20)
		// While theoretically possible to get the same string, it's extremely unlikely
		if str1 == str2 {
			t.Logf("warning: generated two identical strings (very unlikely): %s", str1)
		}
	})

	t.Run("only contains alphabet characters", func(t *testing.T) {
		str := factory.MakeRandomString(1000)
		for _, c := range str {
			if !strings.ContainsRune(stringx.DefaultRandomAlphabet, c) {
				t.Errorf("generated string contains character %c not in default alphabet", c)
			}
		}
	})

	t.Run("empty string on zero length", func(t *testing.T) {
		if factory.MakeRandomString(0) != "" {
			t.Error("expected empty string for zero length")
		}
	})

	t.Run("panics on negative length", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for negative length")
			}
		}()
		factory.MakeRandomString(-1)
	})
}

func TestRandom(t *testing.T) {
	t.Run("generates string of correct length", func(t *testing.T) {
		lengths := []int{1, 5, 10, 50, 100}
		for _, length := range lengths {
			str := stringx.Random(length)
			if len(str) != length {
				t.Errorf("expected length %d, got %d", length, len(str))
			}
		}
	})
}

func TestRandomFactory_Concurrent(t *testing.T) {
	factory := stringx.NewRandomFactory()
	var wg sync.WaitGroup
	numGoroutines := 100
	strLength := 50

	results := make(chan string, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			str := factory.MakeRandomString(strLength)
			results <- str
		}()
	}

	wg.Wait()
	close(results)

	// Verify all strings have correct length
	for str := range results {
		if len(str) != strLength {
			t.Errorf("expected length %d, got %d", strLength, len(str))
		}
	}
}

func TestRandomAlphabet(t *testing.T) {
	t.Run("numeric only", func(t *testing.T) {
		factory := stringx.NewRandomFactory(stringx.RandomAlphabet("0123456789"))
		str := factory.MakeRandomString(100)
		for _, c := range str {
			if c < '0' || c > '9' {
				t.Errorf("expected only digits, got %c", c)
			}
		}
	})

	t.Run("lowercase only", func(t *testing.T) {
		factory := stringx.NewRandomFactory(stringx.RandomAlphabet("abcdefghijklmnopqrstuvwxyz"))
		str := factory.MakeRandomString(100)
		for _, c := range str {
			if c < 'a' || c > 'z' {
				t.Errorf("expected only lowercase letters, got %c", c)
			}
		}
	})

	t.Run("single character", func(t *testing.T) {
		factory := stringx.NewRandomFactory(stringx.RandomAlphabet("x"))
		str := factory.MakeRandomString(10)
		if str != "xxxxxxxxxx" {
			t.Errorf("expected 'xxxxxxxxxx', got %s", str)
		}
	})

	t.Run("unicode alphabet", func(t *testing.T) {
		unicodeAlphabet := "你好世界"
		factory := stringx.NewRandomFactory(stringx.RandomAlphabet(unicodeAlphabet))
		str := factory.MakeRandomString(100)
		// Verify string length is correct (in runes, not bytes)
		if stringx.Len(str) != 100 {
			t.Errorf("expected rune length 100, got %d", stringx.Len(str))
		}
		// Verify all characters are from the alphabet
		for _, c := range str {
			if !strings.ContainsRune(unicodeAlphabet, c) {
				t.Errorf("generated string contains character %c not in alphabet %s", c, unicodeAlphabet)
			}
		}
	})

	t.Run("mixed ascii and unicode alphabet", func(t *testing.T) {
		mixedAlphabet := "abc你好"
		factory := stringx.NewRandomFactory(stringx.RandomAlphabet(mixedAlphabet))
		str := factory.MakeRandomString(100)
		// Verify string length is correct (in runes, not bytes)
		if stringx.Len(str) != 100 {
			t.Errorf("expected rune length 100, got %d", stringx.Len(str))
		}
		// Verify all characters are from the alphabet
		for _, c := range str {
			if !strings.ContainsRune(mixedAlphabet, c) {
				t.Errorf("generated string contains character %c not in alphabet %s", c, mixedAlphabet)
			}
		}
	})
}

func BenchmarkRandomFactory_MakeRandomString(b *testing.B) {
	factory := stringx.NewRandomFactory()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		factory.MakeRandomString(32)
	}
}

func BenchmarkRandomFactory_MakeRandomString_Concurrent(b *testing.B) {
	factory := stringx.NewRandomFactory()
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			factory.MakeRandomString(32)
		}
	})
}
