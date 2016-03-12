package bubblebabble_test

import (
	"fmt"

	"github.com/martinlindhe/bubblebabble"
)

func ExampleEncode() {

	data := []byte("Pineapple")

	dst := make([]byte, bubblebabble.EncodedLen(len(data)))
	bubblebabble.Encode(dst, data)

	fmt.Printf(string(dst))
	// Output: xigak-nyryk-humil-bosek-sonax
}

func ExampleDecode() {

	encoded := []byte("xigak-nyryk-humil-bosek-sonax")

	dst := make([]byte, bubblebabble.MaxDecodedLen(len(encoded)))
	bubblebabble.Decode(dst, encoded)

	fmt.Printf(string(dst))
	// Output: Pineapple
}
