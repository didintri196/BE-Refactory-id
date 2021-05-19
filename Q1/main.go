package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Transaksi struct {
	NamaToko      string
	Tanggal       string
	Namakasir     string
	TransaksiData []TransaksiData
}

type TransaksiData struct {
	Nama  string
	Harga int64
}

func maxtype(data string) (dataprint []string) {
	tmp := ""
	ix := 1
	for _, rowdata := range data {
		if ix%30 == 0 {
			tmp += string(rowdata)
			dataprint = append(dataprint, tmp)
			tmp = ""
		} else {
			tmp += string(rowdata)
		}
		ix++
	}
	if tmp != "" {
		dataprint = append(dataprint, tmp)
	}
	return
}

func dot(nama string, hargaint int64) (hasil string) {
	harga := FormatInt64ToRp(hargaint)
	totnama := len(nama)
	totharga := len(harga)
	totall := totnama + totharga
	mod := totall % 30
	dotx := ""
	if mod < 30 {
		kurang := 30 - mod
		for i := 0; i < kurang; i++ {
			dotx += "."
		}
	}

	hasil = nama + dotx + harga
	return
}

func FormatInt64ToRp(amount int64) string {

	s := fmt.Sprintf("%d", amount)

	r := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != s; {
		n = s
		s = r.ReplaceAllString(s, "$1.$2")
	}

	if amount < 0 {
		s = strings.Replace(s, "-", "", -1)

		return fmt.Sprintf("-Rp%s", s)
	}

	return fmt.Sprintf("Rp%s", s)

}
func input() (transaksi *Transaksi) {
	transaksi = &Transaksi{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("NAMA TOKO-> ")
	namatoko, _ := reader.ReadString('\n')
	transaksi.NamaToko = strings.Replace(namatoko, "\n", "", -1)

	fmt.Print("TANGGAL-> ")
	tanggal, _ := reader.ReadString('\n')
	transaksi.Tanggal = strings.Replace(tanggal, "\n", "", -1)

	fmt.Print("NAMA KASIR-> ")
	namakasir, _ := reader.ReadString('\n')
	transaksi.Namakasir = strings.Replace(namakasir, "\n", "", -1)
	var datatransaksi []TransaksiData
	for {
		fmt.Print("ITEM-> ")
		nama, _ := reader.ReadString('\n')
		nama = strings.Replace(nama, "\n", "", -1)

		fmt.Print("PRICE-> ")
		harga, _ := reader.ReadString('\n')
		harga = strings.Replace(harga, "\n", "", -1)
		hargaint64, _ := strconv.ParseInt(harga, 10, 64)
		datatransaksi = append(datatransaksi, TransaksiData{
			Nama:  nama,
			Harga: hargaint64,
		})

		fmt.Print("input again? (yes/exit)")
		say, _ := reader.ReadString('\n')
		say = strings.Replace(say, "\n", "", -1)
		if strings.Compare("exit", say) == 0 {
			transaksi.TransaksiData = datatransaksi
			break
		}
	}
	return
}
func main() {
	transaksi := input()
	data_print := maxtype(transaksi.NamaToko)
	for _, dataok := range data_print {
		fmt.Println(dataok)
	}
	data_print = maxtype("Tanggal : " + transaksi.Tanggal)
	for _, dataok := range data_print {
		fmt.Println(dataok)
	}
	data_print = maxtype("Kasir   : " + transaksi.Namakasir)
	for _, dataok := range data_print {
		fmt.Println(dataok)
	}
	fmt.Println("==============================")
	total := 0
	for _, rowtrans := range transaksi.TransaksiData {
		data := dot(rowtrans.Nama, rowtrans.Harga)
		data_print := maxtype(data)
		for _, dataok := range data_print {
			fmt.Println(dataok)
		}
		total = total + int(rowtrans.Harga)
	}
	fmt.Println()
	data := dot("Total", int64(total))
	data_print = maxtype(data)
	for _, dataok := range data_print {
		fmt.Println(dataok)
	}
}
