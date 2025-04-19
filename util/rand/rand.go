package rand

import (
	"crypto/rand"
	mathrand "math/rand"
	"time"

	"github.com/rs/xid"
)

var (
	letterRunes       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hexRunes          = []rune("0123456789abcdef")
	alphaNumericRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func init() {
	// This helps generate different numbers on each start, but doesn't make the library secure cryptographically.
	mathrand.Seed(time.Now().UnixNano())
}

// RandomID generates a random string with xid with a prefix.
func RandomID(prefix string) string {
	return prefix + "-" + xid.New().String()
}

// RandomBytes generates a byte slice containing random data of given size.
// It depends on the OS whether it's cryptographically secure.
func RandomBytes(size int) []byte {
	bytes := make([]byte, size)
	_, _ = rand.Reader.Read(bytes)
	return bytes
}

// RandomLetters generates a random string containing letters.
func RandomLetters(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mathrand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomAlphaNumerics(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumericRunes[mathrand.Intn(len(alphaNumericRunes))]
	}
	return string(b)
}

// RandomHex generates a random hex-string.
func RandomHex(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = hexRunes[mathrand.Intn(len(hexRunes))]
	}
	return string(b)
}

// RandomUInt8 generates a random int ranged in [l, r].
func RandomUInt8(l, r uint8) uint8 {
	return uint8(RandomInt(int(l), int(r)))
}

// RandomInt generates a random int ranged in [l, r].
func RandomInt(l, r int) int {
	return l + mathrand.Intn(r-l+1)
}

// RandomInt64 generates a random int64 ranged in [l, r].
func RandomInt64(l, r int64) int64 {
	return l + mathrand.Int63n(r-l+1)
}
