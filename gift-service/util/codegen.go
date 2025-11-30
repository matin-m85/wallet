package util

import (
	"crypto/rand"
	"fmt"
	"strings"
	"unicode"
)

func GeneratePrefix(name string) string {
	var b strings.Builder
	for _, r := range name {
		if len(b.String()) >= 4 {
			break
		}
		ru := unicode.ToUpper(r)
		if ru <= 127 && (unicode.IsLetter(ru) || unicode.IsDigit(ru)) {
			b.WriteRune(ru)
		}
	}
	p := b.String()

	p = strings.Map(func(r rune) rune {
		if (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, p)
	if len(p) >= 4 {
		return p[:4]
	}
	for len(p) < 4 {
		p = p + "X"
	}
	return p
}

func GenerateRandomBody(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		for i := range b {
			b[i] = letters[i%len(letters)]
		}
	} else {
		for i := range b {
			b[i] = letters[int(b[i])%len(letters)]
		}
	}
	return string(b)
}

func GenerateCode(prefix string) string {
	body := GenerateRandomBody(8)
	return fmt.Sprintf("%s-%s", prefix, body)
}
