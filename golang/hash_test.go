package golang

import (
	"bytes"
	"fmt"
	"github.com/spaolacci/murmur3"
	"github.com/ugorji/go/codec"
	"testing"
)

func TestExamples(t *testing.T) {
	var examples = map[string]string{
		`{}`:      "466e20057851c2d220882a78996617be",
		`  {  } `: "466e20057851c2d220882a78996617be",
		`{   "uuid":  "f86e77a4-0dbd-4e13-98b5-c5c97c34611a" }`: "b0720da3fb340cee79559f8bc9233d0d",
		`{"b":"b", "c":"c", "a":"a"}`:                           "4a3ab9c9f7b795e09b6d2a822f5552ba",
		`{"x":1,
		"y":2}`: "4a2f0489c0b8a54e32804edf876a6df4",
	}

	for input, expectedHash := range examples {
		formatted := formatJSON(input)
		t.Logf("formatted json is %s\n", formatted)
		hash := murmurHash(formatted)
		if hash != expectedHash {
			t.Errorf("hash mismatch for '%s'.  expected %s but got %s", input, expectedHash, hash)
		}
	}
}

func formatJSON(jsonInput string) string {

	in := make(map[string]interface{})

	codec.NewDecoderBytes([]byte(jsonInput), new(codec.JsonHandle)).MustDecode(in)

	var buf bytes.Buffer

	encHand := &codec.JsonHandle{}
	encHand.Canonical = true

	enc := codec.NewEncoder(&buf, encHand)
	if err := enc.Encode(in); err != nil {
		panic(err)
	}

	return string(buf.Bytes())
}

func murmurHash(in string) string {
	h := murmur3.New128()
	if _, err := h.Write([]byte(in)); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
