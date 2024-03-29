package bubblebabble

import (
	"bytes"
	"testing"
)

type testpair struct {
	decoded, encoded []byte
}

func makePair(dec, enc string) testpair {
	return testpair{[]byte(dec), []byte(enc)}
}

var pairs = []testpair{
	makePair("", "xexax"),
	makePair("1234567890", "xesef-disof-gytuf-katof-movif-baxux"),
	makePair("Pineapple", "xigak-nyryk-humil-bosek-sonax"),
	makePair("asdf", "ximel-finek-koxex"),
}

func TestEncode(t *testing.T) {
	for i, p := range pairs {
		dst := make([]byte, EncodedLen(len(p.decoded)))
		n := Encode(dst, p.decoded)
		if n != len(dst) {
			t.Errorf("#%d: Encode returned %d, expected %d", i, n, len(dst))
		}
		if !bytes.Equal(dst, p.encoded) {
			t.Errorf("#%d: Encode encoded %#v, expected %#v", i, string(dst), string(p.encoded))
		}
	}
}

func TestDecode(t *testing.T) {
	for i, p := range pairs {
		dst := make([]byte, MaxDecodedLen(len(p.encoded)))
		n, err := Decode(dst, p.encoded)
		dst = dst[0:n]
		if err != nil {
			t.Errorf("#%d: Decoding %#v caused error %#v", i, string(p.encoded), err)
		}
		if !bytes.Equal(dst, p.decoded) {
			t.Errorf("#%d: Decode decoded %#v, expected %#v", i, string(dst), string(p.decoded))
		}
	}
}

type corruptPair struct {
	encoded []byte
	err     error
}

func makeCorrupt(enc string, err error) corruptPair {
	return corruptPair{[]byte(enc), err}
}

var corrupts = []corruptPair{
	makeCorrupt("Ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn", CorruptInputError(0)),
	makeCorrupt("", CorruptInputError(0)),
	makeCorrupt("xexux", CorruptInputError(3)),
	makeCorrupt("nyryk-humil-bosek", CorruptInputError(0)),
	makeCorrupt("xigak-nyryk-humil-bosek", CorruptInputError(22)),
	makeCorrupt("nyryk-humil-bosek-sonax", CorruptInputError(0)),
}

func TestDecodeCorrupt(t *testing.T) {
	for i, corrupt := range corrupts {
		dst := make([]byte, MaxDecodedLen(len(corrupt.encoded)))
		_, err := Decode(dst, corrupt.encoded)
		if err == nil {
			t.Errorf("#%d: Decoder failed to detect corruption in %#v", i, string(corrupt.encoded))
		}
		if err != corrupt.err {
			t.Errorf("#%d: Expected err %#v, got %#v", i, corrupt.err, err)
		}
	}
}
