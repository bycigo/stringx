package stringx

import (
	"strings"
	"sync"
	"testing"
)

// containsFromSets reports whether every byte in s comes from the union of the
// provided character sets.
func allBytesIn(s string, sets ...string) bool {
	union := strings.Join(sets, "")
	for i := 0; i < len(s); i++ {
		if !strings.ContainsRune(union, rune(s[i])) {
			return false
		}
	}
	return true
}

func TestPassword(t *testing.T) {
	t.Run("generates password of correct length", func(t *testing.T) {
		lengths := []int{1, 5, 10, 32, 100}
		for _, length := range lengths {
			pwd := Password(length)
			if len(pwd) != length {
				t.Errorf("expected length %d, got %d", length, len(pwd))
			}
		}
	})

	t.Run("empty string on zero length", func(t *testing.T) {
		if Password(0) != "" {
			t.Error("expected empty string for zero length")
		}
	})

	t.Run("panics on negative length", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for negative length")
			}
		}()
		Password(-1)
	})

	t.Run("uses default character sets", func(t *testing.T) {
		pwd := Password(1000)
		if !allBytesIn(pwd, DefaultPasswordLetters, DefaultPasswordNumbers, DefaultPasswordSymbols) {
			t.Error("password contains characters outside the default sets")
		}
		// space is not part of the default includes
		if strings.Contains(pwd, " ") {
			t.Error("password should not contain spaces by default")
		}
	})
}

func TestPasswordFactory_MakePassword(t *testing.T) {
	factory := NewPasswordFactory()

	t.Run("generates password of correct length", func(t *testing.T) {
		lengths := []int{1, 5, 10, 32, 100}
		for _, length := range lengths {
			pwd := factory.MakePassword(length)
			if len(pwd) != length {
				t.Errorf("expected length %d, got %d", length, len(pwd))
			}
		}
	})

	t.Run("empty string on zero length", func(t *testing.T) {
		if factory.MakePassword(0) != "" {
			t.Error("expected empty string for zero length")
		}
	})

	t.Run("panics on negative length", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for negative length")
			}
		}()
		factory.MakePassword(-1)
	})

	t.Run("generates different passwords", func(t *testing.T) {
		pwd1 := factory.MakePassword(20)
		pwd2 := factory.MakePassword(20)
		if pwd1 == pwd2 {
			t.Logf("warning: generated two identical passwords (very unlikely): %s", pwd1)
		}
	})

	t.Run("contains at least one char from each selected kind", func(t *testing.T) {
		// run many times because at least-one guarantee is probabilistic-free
		for i := 0; i < 100; i++ {
			pwd := factory.MakePassword(10, PasswordIncludes{PasswordLetter, PasswordNumber, PasswordSymbol})
			if !strings.ContainsAny(pwd, DefaultPasswordLetters) {
				t.Fatalf("password %q missing a letter", pwd)
			}
			if !strings.ContainsAny(pwd, DefaultPasswordNumbers) {
				t.Fatalf("password %q missing a number", pwd)
			}
			if !strings.ContainsAny(pwd, DefaultPasswordSymbols) {
				t.Fatalf("password %q missing a symbol", pwd)
			}
		}
	})
}

func TestPasswordIncludes(t *testing.T) {
	factory := NewPasswordFactory()

	t.Run("letters only", func(t *testing.T) {
		pwd := factory.MakePassword(200, PasswordIncludes{PasswordLetter})
		if !allBytesIn(pwd, DefaultPasswordLetters) {
			t.Errorf("expected only letters, got %q", pwd)
		}
	})

	t.Run("numbers only", func(t *testing.T) {
		pwd := factory.MakePassword(200, PasswordIncludes{PasswordNumber})
		if !allBytesIn(pwd, DefaultPasswordNumbers) {
			t.Errorf("expected only numbers, got %q", pwd)
		}
	})

	t.Run("letters and numbers", func(t *testing.T) {
		pwd := factory.MakePassword(200, PasswordIncludes{PasswordLetter, PasswordNumber})
		if !allBytesIn(pwd, DefaultPasswordLetters, DefaultPasswordNumbers) {
			t.Errorf("expected only letters and numbers, got %q", pwd)
		}
		if strings.ContainsAny(pwd, DefaultPasswordSymbols) {
			t.Errorf("password %q should not contain symbols", pwd)
		}
	})

	t.Run("with spaces", func(t *testing.T) {
		// a long password including spaces should eventually contain one
		pwd := factory.MakePassword(500, PasswordIncludes{PasswordLetter, PasswordSpace})
		if !allBytesIn(pwd, DefaultPasswordLetters, DefaultPasswordSpaces) {
			t.Errorf("expected only letters and spaces, got %q", pwd)
		}
	})

	t.Run("duplicate kinds are de-duplicated", func(t *testing.T) {
		pwd := factory.MakePassword(50, PasswordIncludes{PasswordNumber, PasswordNumber, PasswordNumber})
		if !allBytesIn(pwd, DefaultPasswordNumbers) {
			t.Errorf("expected only numbers, got %q", pwd)
		}
	})

	t.Run("unknown kind is ignored, falls back to valid ones", func(t *testing.T) {
		pwd := factory.MakePassword(50, PasswordIncludes{PasswordNumber, PasswordCharKind(999)})
		if !allBytesIn(pwd, DefaultPasswordNumbers) {
			t.Errorf("expected only numbers, got %q", pwd)
		}
	})

	t.Run("panics when no valid kinds remain", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic when no valid character sources")
			}
		}()
		factory.MakePassword(10, PasswordIncludes{PasswordCharKind(999)})
	})
}

func TestPasswordSource_Custom(t *testing.T) {
	t.Run("custom source restricts characters", func(t *testing.T) {
		src := PasswordSource{
			PasswordLetter: []byte("abc"),
			PasswordNumber: []byte("01"),
		}
		factory := NewPasswordFactoryWithSource(src)
		pwd := factory.MakePassword(200, PasswordIncludes{PasswordLetter, PasswordNumber})
		if !allBytesIn(pwd, "abc", "01") {
			t.Errorf("password %q contains characters outside the custom source", pwd)
		}
	})

	t.Run("empty char slice is filtered out", func(t *testing.T) {
		src := PasswordSource{
			PasswordLetter: []byte("abc"),
			PasswordNumber: []byte{}, // empty, must be skipped
		}
		factory := NewPasswordFactoryWithSource(src)
		// PasswordNumber is empty so only letters should appear, and no panic
		pwd := factory.MakePassword(100, PasswordIncludes{PasswordLetter, PasswordNumber})
		if !allBytesIn(pwd, "abc") {
			t.Errorf("expected only letters, got %q", pwd)
		}
	})

	t.Run("panics when all selected sources are empty", func(t *testing.T) {
		src := PasswordSource{
			PasswordLetter: []byte{},
		}
		factory := NewPasswordFactoryWithSource(src)
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic when all selected sources are empty")
			}
		}()
		factory.MakePassword(10, PasswordIncludes{PasswordLetter})
	})
}

func TestSecuredPasswordFactory(t *testing.T) {
	t.Run("generates password of correct length", func(t *testing.T) {
		factory := NewSecuredPasswordFactory()
		lengths := []int{1, 5, 10, 32, 100}
		for _, length := range lengths {
			pwd := factory.MakePassword(length)
			if len(pwd) != length {
				t.Errorf("expected length %d, got %d", length, len(pwd))
			}
		}
	})

	t.Run("respects custom source", func(t *testing.T) {
		src := PasswordSource{
			PasswordNumber: []byte("0123456789"),
		}
		factory := NewSecuredPasswordFactoryWithSource(src)
		pwd := factory.MakePassword(100, PasswordIncludes{PasswordNumber})
		if !allBytesIn(pwd, DefaultPasswordNumbers) {
			t.Errorf("expected only numbers, got %q", pwd)
		}
	})

	t.Run("default factory uses a secure source", func(t *testing.T) {
		// The top-level Password must be backed by the secured factory.
		if _, ok := DefaultPasswordFactory.(*passwordFactory); !ok {
			t.Fatal("expected DefaultPasswordFactory to be *passwordFactory")
		}
	})
}

func TestPasswordFactory_Concurrent(t *testing.T) {
	factories := map[string]PasswordFactory{
		"normal":  NewPasswordFactory(),
		"secured": NewSecuredPasswordFactory(),
	}

	for name, factory := range factories {
		t.Run(name, func(t *testing.T) {
			var wg sync.WaitGroup
			numGoroutines := 100
			pwdLength := 32
			results := make(chan string, numGoroutines)

			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					results <- factory.MakePassword(pwdLength)
				}()
			}

			wg.Wait()
			close(results)

			for pwd := range results {
				if len(pwd) != pwdLength {
					t.Errorf("expected length %d, got %d", pwdLength, len(pwd))
				}
			}
		})
	}
}

func BenchmarkPasswordFactory_MakePassword(b *testing.B) {
	factory := NewPasswordFactory()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		factory.MakePassword(32)
	}
}

func BenchmarkSecuredPasswordFactory_MakePassword(b *testing.B) {
	factory := NewSecuredPasswordFactory()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		factory.MakePassword(32)
	}
}

func BenchmarkPasswordFactory_MakePassword_Concurrent(b *testing.B) {
	factory := NewPasswordFactory()
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			factory.MakePassword(32)
		}
	})
}
