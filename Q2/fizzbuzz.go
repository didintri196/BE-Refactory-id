package q2

import (
	"fmt"
	"strconv"
)

func CekFizzBuzz(i int) string {
	fizz := "Fizz"
	buzz := "Buzz"

	if i%3 == 0 && i%5 == 0 {
		return fizz + buzz
	} else if i%3 == 0 {
		return fizz
	} else if i%5 == 0 {
		return buzz
	} else {
		return strconv.Itoa(i)
	}
}

func FizzBuzz(n int) {
	var hasil []string
	for i := 1; i <= n; i++ {
		hasil = append(hasil, CekFizzBuzz(i))
	}
	fmt.Println(hasil)
}
