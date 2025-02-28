package main

import "fmt"

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// var sample = [8]byte{'\xbd', '\xb2', '\x3d', '\xbc', '\x20', '\xe2', '\x8c', '\x98'}
	// var sample = []byte{'\xbd', '\xb2', '\x3d', '\xbc', '\x20', '\xe2', '\x8c', '\x98'}

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		// fmt.Printf("%x %b %s\n", sample[i], sample[i], string(sample[i]))
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	const nihongo = "日本語"
	for i := 0; i < len(nihongo); i++ {
		fmt.Printf("%#U", nihongo[i])
	}
}
