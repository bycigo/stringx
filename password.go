package stringx

import (
	crand "crypto/rand"
	"math/big"
	"math/rand/v2"
)

// Password generates a password of the specified length.
func Password(length int, opt ...PasswordOption) string {
	return DefaultPasswordFactory.MakePassword(length, opt...)
}

// PasswordFactory is an interface for generating passwords.
type PasswordFactory interface {
	// MakePassword generates a password of the specified length.
	MakePassword(length int, opt ...PasswordOption) string
}

// PasswordOption is an interface for applying password options.
type PasswordOption interface {
	// ApplyPasswordOptions applies the option to the given PasswordOptions.
	ApplyPasswordOptions(opts *PasswordOptions)
}

// PasswordOptions contains options for generating passwords.
type PasswordOptions struct {
	// Includes specifies which character kinds to include in the password.
	Includes PasswordIncludes
}

// Join applies the given PasswordOption to the PasswordOptions.
func (o *PasswordOptions) Join(opts ...PasswordOption) {
	for _, opt := range opts {
		opt.ApplyPasswordOptions(o)
	}
}

// PasswordIncludes is a PasswordOption that specifies which character kinds to include in the password.
type PasswordIncludes []PasswordCharKind

// ApplyPasswordOptions applies the PasswordIncludes option to the given PasswordOptions.
func (o PasswordIncludes) ApplyPasswordOptions(opts *PasswordOptions) {
	opts.Includes = o
}

// NewPasswordFactory creates a new PasswordFactory.
func NewPasswordFactory() PasswordFactory {
	return newPasswordFactory(newDefaultPasswordSource())
}

// NewPasswordFactoryWithSource creates a new PasswordFactory with the given PasswordSource.
func NewPasswordFactoryWithSource(src PasswordSource) PasswordFactory {
	return newPasswordFactory(src)
}

// NewSecuredPasswordFactory creates a new PasswordFactory
// that uses crypto/rand for secure random number generation.
func NewSecuredPasswordFactory() PasswordFactory {
	return newSecuredPasswordFactory(newDefaultPasswordSource())
}

// NewSecuredPasswordFactoryWithSource creates a new PasswordFactory
// that uses crypto/rand for secure random number generation with the given PasswordSource.
func NewSecuredPasswordFactoryWithSource(src PasswordSource) PasswordFactory {
	return newSecuredPasswordFactory(src)
}

// PasswordCharKind represents the kind of characters to include in the password.
type PasswordCharKind int

const (
	_ PasswordCharKind = iota
	PasswordLetter
	PasswordNumber
	PasswordSymbol
	PasswordSpace
)

// PasswordSource is a map of character kinds to their corresponding byte slices.
type PasswordSource map[PasswordCharKind][]byte

func (s PasswordSource) filterCharKinds(opts *PasswordOptions) []PasswordCharKind {
	kinds := make([]PasswordCharKind, 0)
	set := make(map[PasswordCharKind]struct{})
	for _, kind := range opts.Includes {
		if _, ok := set[kind]; ok {
			continue
		}
		if chars, ok := s[kind]; ok && len(chars) > 0 {
			kinds = append(kinds, kind)
			set[kind] = struct{}{}
		}
	}
	return kinds
}

// DefaultPasswordFactory is the default PasswordFactory used for generating passwords.
// It uses crypto/rand for secure random number generation, so the top-level Password
// function is safe for generating real passwords by default.
var DefaultPasswordFactory PasswordFactory = NewSecuredPasswordFactory()

const (
	// DefaultPasswordLetters is the default set of letters used for generating passwords.
	DefaultPasswordLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// DefaultPasswordNumbers is the default set of numbers used for generating passwords.
	DefaultPasswordNumbers = "0123456789"
	// DefaultPasswordSymbols is the default set of symbols used for generating passwords.
	DefaultPasswordSymbols = "~!@#$%^&*()-_=+[]{}|;:,.<>?/"
	// DefaultPasswordSpaces is the default set of spaces used for generating passwords.
	DefaultPasswordSpaces = " "
)

func newDefaultPasswordIncludes() PasswordIncludes {
	return PasswordIncludes{PasswordLetter, PasswordNumber, PasswordSymbol}
}

func newDefaultPasswordSource() PasswordSource {
	return PasswordSource{
		PasswordLetter: []byte(DefaultPasswordLetters),
		PasswordNumber: []byte(DefaultPasswordNumbers),
		PasswordSymbol: []byte(DefaultPasswordSymbols),
		PasswordSpace:  []byte(DefaultPasswordSpaces),
	}
}

type passwordFactory struct {
	source   PasswordSource
	randIntN func(n int) int
}

// MakePassword generates a password of the specified length.
func (f *passwordFactory) MakePassword(length int, opt ...PasswordOption) string {
	if length < 0 {
		panic("negative password length")
	}
	if length == 0 {
		return ""
	}

	opts := PasswordOptions{
		Includes: newDefaultPasswordIncludes(),
	}
	opts.Join(opt...)

	kinds := f.source.filterCharKinds(&opts)
	if len(kinds) == 0 {
		panic("no valid character sources for password generation")
	}

	password := make([]byte, 0, length)

	// ensure that the password contains at least one character from each selected
	// source, while never exceeding the requested length
	for _, kind := range kinds {
		if len(password) >= length {
			break
		}
		password = append(password, f.randByte(f.source[kind]))
	}

	// fill the remaining length by picking a random source per character, which
	// keeps the distribution across sources uniform
	for len(password) < length {
		kind := kinds[f.randIntN(len(kinds))]
		password = append(password, f.randByte(f.source[kind]))
	}

	f.shuffle(password)

	return FromBytes(password)
}

func (f *passwordFactory) randByte(bs []byte) byte {
	return bs[f.randIntN(len(bs))]
}

func (f *passwordFactory) shuffle(bs []byte) {
	for i := len(bs) - 1; i > 0; i-- {
		j := f.randIntN(i + 1)
		bs[i], bs[j] = bs[j], bs[i]
	}
}

func newPasswordFactory(src PasswordSource) PasswordFactory {
	return &passwordFactory{
		source:   src,
		randIntN: rand.IntN,
	}
}

func newSecuredPasswordFactory(src PasswordSource) PasswordFactory {
	return &passwordFactory{
		source: src,
		randIntN: func(n int) int {
			v, err := crand.Int(crand.Reader, big.NewInt(int64(n)))
			if err != nil {
				panic("failed to read secure random number: " + err.Error())
			}
			return int(v.Int64())
		},
	}
}
