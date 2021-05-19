package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func randomstring(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}

func post(i string) {

	url := "http://localhost:9999"

	payload := strings.NewReader("{\"counter\":\"" + i + "\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-RANDOM", randomstring(8))
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "61f4dac2-91f3-4ddb-94e7-00df46dbccce")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func main() {
	i := 1
	// go func() {
	for {
		post(strconv.Itoa(i))
		i++
		time.Sleep(1 * time.Second)
	}
	// }()
}
