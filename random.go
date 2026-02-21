package stringx

import (
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// DefaultRandomFactory is the default RandomFactory used for generating random strings.
var DefaultRandomFactory = NewRandomFactory()

// RandomFactory is an interface for generating random strings.
type RandomFactory interface {
	// MakeRandomString generates a random string of the specified length.
	MakeRandomString(length int) string
}

// DefaultRandomAlphabet is the default alphabet used for generating random strings.
const DefaultRandomAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomOptions contains options for generating random strings.
type RandomOptions struct {
	// Alphabet is the set of characters to use for generating random strings.
	Alphabet []rune
}

// RandomOption is a function that modifies RandomOptions.
type RandomOption func(opts *RandomOptions)

// RandomAlphabet sets the alphabet for generating random strings.
func RandomAlphabet(alphabet string) RandomOption {
	return func(opts *RandomOptions) {
		opts.Alphabet = []rune(alphabet)
	}
}

var randomSourcePool = newRandomSourceFactory()

type randomSourceFactory struct {
	seedCount  atomic.Int64
	sourcePool sync.Pool
}

func newRandomSourceFactory() *randomSourceFactory {
	f := &randomSourceFactory{}
	f.sourcePool = sync.Pool{
		New: func() any {
			count := f.seedCount.Add(1)
			seed := time.Now().UnixNano() + count
			return rand.NewSource(seed)
		},
	}
	return f
}

func (f *randomSourceFactory) acquireSource() rand.Source {
	return f.sourcePool.Get().(rand.Source)
}

func (f *randomSourceFactory) releaseSource(source rand.Source) {
	f.sourcePool.Put(source)
}

type randomFactory struct {
	opts RandomOptions
}

// NewRandomFactory creates a new RandomFactory.
func NewRandomFactory(opt ...RandomOption) RandomFactory {
	opts := RandomOptions{
		Alphabet: []rune(DefaultRandomAlphabet),
	}
	for _, o := range opt {
		o(&opts)
	}
	return &randomFactory{
		opts: opts,
	}
}

// MakeRandomString generates a random string of the specified length.
func (f *randomFactory) MakeRandomString(length int) string {
	if length < 0 {
		panic("negative random string length")
	}
	if length == 0 {
		return ""
	}

	source := randomSourcePool.acquireSource()
	defer randomSourcePool.releaseSource(source)

	alphabetRunes := f.opts.Alphabet
	alphabetLen := int64(len(alphabetRunes))

	b := strings.Builder{}
	b.Grow(length)
	for i := 0; i < length; i++ {
		b.WriteRune(alphabetRunes[source.Int63()%alphabetLen])
	}
	return b.String()
}

// Random generates a random string of the specified length.
func Random(length int) string {
	return DefaultRandomFactory.MakeRandomString(length)
}
