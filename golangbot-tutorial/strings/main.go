package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Basic strings
	name := "Hello world"
	fmt.Println(name)
	printBytes(name)
	fmt.Println()
	printByteChars(name)
	fmt.Println("---")

	// Multi-byte chars and runes
	specialChar := "Señor"
	fmt.Println("length", len(specialChar))
	printByteChars(specialChar)
	fmt.Println()
	printChars(specialChar)
	fmt.Println()
	printRunes(specialChar)
	fmt.Println("---")

	// Constructing strings
	byteSliceHex := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	strHex := string(byteSliceHex)
	fmt.Println(strHex)

	byteSliceDec := []byte{67, 97, 102, 195, 169}
	strDec := string(byteSliceDec)
	fmt.Println(strDec)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	strRune := string(runeSlice)
	fmt.Println(strRune)

	// String length
	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)

	// String are immutable
	h := "hello"
	fmt.Println(mutate(h))
}

func mutate(s string) string {
	r := []rune(s)
	r[0] = 'a'
	return string(r)
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func printBytes(s string) {
	for i := range s {
		fmt.Printf("%x ", s[i])
	}
}

func printByteChars(s string) {
	for i := range s {
		fmt.Printf("%c ", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := range runes {
		fmt.Printf("%c ", runes[i])
	}
}

func printRunes(s string) {
	for i, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, i)
	}
}
