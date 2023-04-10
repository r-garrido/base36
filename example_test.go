package base36_test

import (
	"fmt"

	"github.com/martinlindhe/base36"
)

func ExampleEncode() {
	base36.SetEncodingAlphabet(base36.Uppercase)
	fmt.Println(base36.Encode(5481594952936519619))
	// Output: 15N9Z8L3AU4EB
}

func ExampleDecode() {
	base36.SetEncodingAlphabet(base36.Uppercase)
	fmt.Println(base36.Decode("15N9Z8L3AU4EB"))
	// Output: 5481594952936519619
}

func ExampleEncodeBytes() {
	base36.SetEncodingAlphabet(base36.Uppercase)
	fmt.Println(base36.EncodeBytes([]byte{1, 2, 3, 4}))
	// Output: A2F44
}

func ExampleDecodeToBytes() {
	base36.SetEncodingAlphabet(base36.Uppercase)
	fmt.Println(base36.DecodeToBytes("A2F44"))
	// Output: [1 2 3 4]
}

func ExampleEncodeLower() {
	base36.SetEncodingAlphabet(base36.Lowercase)
	fmt.Println(base36.Encode(5481594952936519619))
	// Output: 15n9z8l3au4eb
}

func ExampleDecodeLower() {
	base36.SetEncodingAlphabet(base36.Lowercase)
	fmt.Println(base36.Decode("15n9z8l3au4eb"))
	// Output: 5481594952936519619
}

func ExampleEncodeLowerBytes() {
	base36.SetEncodingAlphabet(base36.Lowercase)
	fmt.Println(base36.EncodeBytes([]byte{1, 2, 3, 4}))
	// Output: a2f44
}

func ExampleDecodeLowerToBytes() {
	base36.SetEncodingAlphabet(base36.Lowercase)
	fmt.Println(base36.DecodeToBytes("a2f44"))
	// Output: [1 2 3 4]
}
