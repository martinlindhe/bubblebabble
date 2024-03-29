package bubblebabble_test

import (
	"fmt"

	"github.com/martinlindhe/bubblebabble"
)

func ExampleEncode() {
	data := []byte("Pineapple")

	dst := make([]byte, bubblebabble.EncodedLen(len(data)))
	bubblebabble.Encode(dst, data)

	fmt.Println(string(dst))
	// Output: xigak-nyryk-humil-bosek-sonax
}

func ExampleEncodeToString() {
	data := []byte("Pineapple")

	dst := bubblebabble.EncodeToString(data)

	fmt.Println(dst)
	// Output: xigak-nyryk-humil-bosek-sonax
}

func ExampleDecode() {
	encoded := []byte("xigak-nyryk-humil-bosek-sonax")

	dst := make([]byte, bubblebabble.MaxDecodedLen(len(encoded)))

	n, _ := bubblebabble.Decode(dst, encoded)
	dst = dst[0:n]

	fmt.Println(string(dst))
	// Output: Pineapple
}

func ExampleDecodeString() {
	dst, _ := bubblebabble.DecodeString("xigak-nyryk-humil-bosek-sonax")

	fmt.Println(string(dst))
	// Output: Pineapple
}
