package q2

import (
	"fmt"
	"testing"
)

func TestLapyear(t *testing.T) {
	fmt.Println("[Test Lapyear]")
	awal := 1900
	akhir := 2020
	for i := awal; i <= akhir; i++ {
		if CekLapyear(i) {
			fmt.Print(i, ",")
		}
	}
	fmt.Println()
}
