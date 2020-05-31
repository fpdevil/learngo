package main

import (
	"fmt"
	"unicode/utf8"
)

// a text from the hitchhikers guide to galaxy in Turkish
const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

func main() {
	fmt.Println("--- [01] 04-rune-decoding ---")

	r, size := utf8.DecodeRuneInString("öykü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ykü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("kü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	// the for loop cannot print the multi byte runes,
	// instead it prints bytes
	for i := 0; i < len(text); i++ {
		fmt.Printf("%c", text[i])
	}
	fmt.Println()

	fmt.Println("\n\n\n")
	// printing runes instead of bytes is decoding
	fmt.Println("=== attempting multi byte decoding with runes ===")
	for i := 0; i < len(text); i++ {
		// this converts a byte to a rune, but does not
		// solve the problem
		r := rune(text[i])
		fmt.Printf("%c", r)
	}

	fmt.Println("\n\n\n")
	fmt.Println("=== again attempting multi byte decoding with runes ===")
	// checking where the rune starts and ends using range
	for i := 0; i < len(text); {
		// r, _ := utf8.DecodeRuneInString(text)
		// fmt.Printf("%c", r)
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)
		i += size
	}

	// a more efficient approach
	fmt.Println("\n\n\n")
	fmt.Println("=== automatic decoding of runes with for range ===")
	for _, r := range text {
		fmt.Printf("%c", r)
	}
	fmt.Println()
}
