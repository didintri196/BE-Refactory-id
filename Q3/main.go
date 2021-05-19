package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Data struct {
	ID       int        `json:"id"`
	Username string     `json:"username"`
	Profile  Profile    `json:"profile"`
	Articles []Articles `json:"articles:"`
}
type Profile struct {
	FullName string   `json:"full_name"`
	Birthday string   `json:"birthday"`
	Phones   []string `json:"phones"`
}
type Articles struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	PublishedAt string `json:"published_at"`
}

func main() {
	dat, err := ioutil.ReadFile("profile_list.json")
	var datajson []Data
	err = json.Unmarshal(dat, &datajson)

	if err != nil {
		panic(err)
	}

	fmt.Println("normal=>", len(datajson))
	number1 := getdonthavenumber(datajson)
	fmt.Println("getdonthavenumber=>", len(number1))
	number2 := gethavearticles(datajson)
	fmt.Println("gethavearticles=>", len(number2))
	number3 := gethaveannisname(datajson)
	fmt.Println("gethaveannisname=>", len(number3))
	number4 := getarticleon2020(datajson)
	fmt.Println("getarticleon2020=>", len(number4))
	number5 := getborn1986(datajson)
	fmt.Println("getborn1986=>", len(number5))
	number6 := getarticlecontainstips(datajson)
	fmt.Println("getarticlecontainstips=>", len(number6))
	number7 := getarticlepublishedbeforeaugust2019(datajson)
	fmt.Println("getarticlepublishedbeforeaugust2019=>", len(number7))

}

func getdonthavenumber(data []Data) (result []Data) {
	for _, rowdata := range data {
		if len(rowdata.Profile.Phones) == 0 {
			result = append(result, rowdata)
		}
	}
	return
}

func gethavearticles(data []Data) (result []Data) {
	for _, rowdata := range data {
		if len(rowdata.Articles) > 0 {
			result = append(result, rowdata)
		}
	}
	return
}

func gethaveannisname(data []Data) (result []Data) {
	for _, rowdata := range data {
		if strings.Contains(strings.ToLower(rowdata.Profile.FullName), "annis") {
			result = append(result, rowdata)
		}
	}
	return
}

func getarticleon2020(data []Data) (result []Data) {
	for _, rowdata := range data {
		for _, rowarticle := range rowdata.Articles {
			var date, _ = time.Parse("2006-01-02T15:04:05", rowarticle.PublishedAt)
			if date.Year() == 2020 {
				result = append(result, rowdata)
				break
			}
		}
	}
	return
}

func getborn1986(data []Data) (result []Data) {
	for _, rowdata := range data {
		var date, _ = time.Parse("2006-01-02", rowdata.Profile.Birthday)
		if date.Year() == 1986 {
			result = append(result, rowdata)
			break
		}
	}
	return
}

func getarticlecontainstips(data []Data) (result []Articles) {
	for _, rowdata := range data {
		for _, rowarticle := range rowdata.Articles {
			if strings.Contains(rowarticle.Title, "Tips") {
				result = append(result, rowarticle)
			}
		}
	}
	return
}

func getarticlepublishedbeforeaugust2019(data []Data) (result []Articles) {
	for _, rowdata := range data {
		for _, rowarticle := range rowdata.Articles {
			var date, _ = time.Parse("2006-01-02T15:04:05", rowarticle.PublishedAt)
			if date.Unix() < 1564617600 {
				result = append(result, rowarticle)
			}
		}
	}
	return
}
