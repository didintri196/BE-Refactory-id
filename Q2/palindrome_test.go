package q2

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	fmt.Println("[Palindrome Test]")
	fmt.Println("Radar =>", Palindrome("Radar"))
	fmt.Println("Malam =>", Palindrome("Malam"))
	fmt.Println("Kasur ini rusak =>", Palindrome("Kasur ini rusak"))
	fmt.Println("Ibu Ratna antar ubi =>", Palindrome("Ibu Ratna antar ubi"))

	fmt.Println("Malas =>", Palindrome("Malas"))
	fmt.Println("Makan nasi goreng =>", Palindrome("Makan nasi goreng"))
	fmt.Println("Balonku ada lima =>", Palindrome("Balonku ada lima"))
}
