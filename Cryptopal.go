package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"unicode/utf8"
)

func HexToBase64(hs string) (string, error) {
	v, err := hex.DecodeString(hs)
	if err != nil {
		return "", err
	}
	log.Printf("%s", v)
	return base64.StdEncoding.EncodeToString(v), nil
}

func StringsToXOR(s1, s2 []byte) []byte {
	if len(s1) != len(s2) {
		panic("xor conversion: mismatch lengths")
	}
	res := make([]byte, len(s1))
	for i := range s1 {
		res[i] = s1[i] ^ s2[i]
	}
	return res
}

func buildCorpus(text string) map[rune]float64 {
	c := make(map[rune]float64)
	for _, char := range text {
		c[char]++
	}
	total := utf8.RuneCountInString(text)
	for char := range c {
		c[char] = c[char] / float64(total)
	}
	return c
}

func scoreEnglish(s string, m map[rune]float64) float64 {
	var score float64
	for _, carac := range s {
		score += m[carac]
	}
	return score / float64(utf8.RuneCountInString(s))
}

func singleXOR(in []byte, key byte) []byte {
	ret := make([]byte, len(in))
	for i, c := range in {
		ret[i] = c ^ key
	}
	return ret
}

func findSingleXORKey(in []byte, c map[rune]float64) []byte {
	var res []byte
	var lastScore float64
	for key := byte(0); key < 255; key++ {
		out := singleXOR(in, key)
		score := scoreEnglish(string(out), c)
		if score > lastScore {
			res = out
			lastScore = score
		}
	}
	return res
}
