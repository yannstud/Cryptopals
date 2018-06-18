package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"math"
	"math/bits"
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

func findSingleXORKey(in []byte, c map[rune]float64) (res []byte, key byte, score float64) {
	for k := byte(0); k < 255; k++ {
		out := singleXOR(in, byte(k))
		s := scoreEnglish(string(out), c)
		if s > score {
			res = out
			key = byte(k)
			score = s
		}
	}
	return
}

func repeatingXOR(in, key []byte) []byte {
	ret := make([]byte, len(in))
	for i := range in {
		ret[i] = in[i] ^ key[i%len(key)]
	}
	return ret
}

func hammingDistance(s1, s2 []byte) int {
	if len(s1) != len(s2) {
		panic("hamming distance: different lenghts")
	}
	var res int
	for i := range s1 {
		res += bits.OnesCount8(s1[i] ^ s2[i])
	}
	return res
}

func findRepeatingXORKeySize(in []byte) int {
	var res int
	bestScore := math.MaxFloat64
	for keyLen := 2; keyLen < 42; keyLen++ {
		a, b := in[:keyLen*3], in[keyLen*3:keyLen*3*2]
		score := float64(hammingDistance(a, b)) / float64(keyLen*10)
		if score < bestScore {
			res = keyLen
			bestScore = score
		}
	}
	return res
}

func findRepeatingXORKey(in []byte, c map[rune]float64) []byte {
	keySize := findRepeatingXORKeySize(in)
	column := make([]byte, (len(in)+keySize-1)/keySize)
	key := make([]byte, keySize)
	for col := 0; col < keySize; col++ {
		for row := range column {
			if row*keySize+col >= len(in) {
				continue
			}
			column[row] = in[row*keySize+col]
		}
		_, k, _ := findSingleXORKey(column, c)
		key[col] = k
	}
	return key
}
