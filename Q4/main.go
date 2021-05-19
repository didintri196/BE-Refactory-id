package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Data struct {
	InventoryID int       `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt int64     `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}
type Placement struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
}

func main() {
	dat, err := ioutil.ReadFile("inventory_list.json")
	var datajson []Data
	err = json.Unmarshal(dat, &datajson)

	if err != nil {
		panic(err)
	}
	// fmt.Println(datajson)
	fmt.Println("normal=>", len(datajson))
	number1 := getinmeetingrooms(datajson)
	fmt.Println("getinmeetingrooms=>", len(number1))
	number2 := getelectronicdevice(datajson)
	fmt.Println("getelectronicdevice=>", len(number2))
	number3 := getfurniture(datajson)
	fmt.Println("getfurniture=>", len(number3))
	number4 := getpurchaseon16jan2020(datajson)
	fmt.Println("getpurchaseon16jan2020=>", len(number4))
	number5 := gettagsbrown(datajson)
	fmt.Println("gettagsbrown=>", len(number5))
}

func getinmeetingrooms(data []Data) (result []Data) {
	for _, rowdata := range data {
		if rowdata.Placement.RoomID == 3 {
			result = append(result, rowdata)
		}
	}
	return
}

func getelectronicdevice(data []Data) (result []Data) {
	for _, rowdata := range data {
		if rowdata.Type == "electronic" {
			result = append(result, rowdata)
		}
	}
	return
}

func getfurniture(data []Data) (result []Data) {
	for _, rowdata := range data {
		if rowdata.Type == "furniture" {
			result = append(result, rowdata)
		}
	}
	return
}

func getpurchaseon16jan2020(data []Data) (result []Data) {
	for _, rowdata := range data {
		var date = time.Unix(rowdata.PurchasedAt, 0)
		if date.Day() == 16 && int(date.Month()) == 1 && date.Year() == 2020 {
			result = append(result, rowdata)
		}
	}
	return
}

func gettagsbrown(data []Data) (result []Data) {
	for _, rowdata := range data {
		for _, rowtags := range rowdata.Tags {
			if rowtags == "brown" {
				result = append(result, rowdata)
			}
		}
	}
	return
}
