package cryptopals

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestProblem1(t *testing.T) {
	res, err := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Fatal(err)
	}
	if res != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Fatal("wrong string", res)
	}

}

func TestProbleme2(t *testing.T) {
	res := StringsToXOR(hexDecode(t, "1c0111001f010100061a024b53535009181c"), hexDecode(t, "686974207468652062756c6c277320657965"))
	if !bytes.Equal(res, hexDecode(t, "746865206b696420646f6e277420706c6179")) {
		t.Errorf("Wrong string %x", res)
	}
}

func readFile(t *testing.T, name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatal("fail to read file:", err)
	}
	return data
}

func base64Decode(t *testing.T, s string) []byte {
	res, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		t.Fatal("fail to decode base64:", s)
	}
	return res
}

func hexDecode(t *testing.T, s string) []byte {
	v, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal("fail to decode hex:", s)
	}
	return v
}

func corpusFromFile(name string) map[rune]float64 {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Sprintln("cant open file:", err))
	}
	return buildCorpus(string(text))
}

var corpus = corpusFromFile("testdata/sherlock.txt")

func TestProbleme3(t *testing.T) {
	res, _, _ := findSingleXORKey(hexDecode(t, "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"), corpus)
	t.Logf("%s", res)
}

func TestProbleme4(t *testing.T) {
	text := readFile(t, "testdata/4.txt")
	var lastScore float64
	var res []byte
	for _, line := range strings.Split(string(text), "\n") {
		out, _, score := findSingleXORKey(hexDecode(t, line), corpus)
		if score > lastScore {
			res = out
			lastScore = score
		}
	}
	t.Logf("%s", res)
}
func TestProbleme5(t *testing.T) {
	input := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	res := repeatingXOR(input, []byte("ICE"))
	if !bytes.Equal(res, hexDecode(t, "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")) {
		t.Error("Wrong baby", res)
	}
}

func TestProbleme6(t *testing.T) {
	if res := hammingDistance([]byte("this is a test"), []byte("wokka wokka!!!")); res != 37 {
		t.Fatal("Wrong hamming distance:", res)
	}
	text := base64Decode(t, string(readFile(t, "testdata/6.txt")))
	t.Log("likely size:", findRepeatingXORKeySize(text))

	key := findRepeatingXORKey(text, corpus)
	t.Logf("likely key: %q", key)

	t.Logf("%s", repeatingXOR(text, key))

}
