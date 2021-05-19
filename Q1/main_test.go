package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	transaksi := Transaksi{
		NamaToko:  "Warung Makan Sederhana",
		Tanggal:   "2020/12/2020 15:30",
		Namakasir: "Didin",
		TransaksiData: []TransaksiData{
			TransaksiData{Nama: "Nasi", Harga: 5000},
			TransaksiData{Nama: "Lauk", Harga: 15000},
			TransaksiData{Nama: "Minum", Harga: 5000},
			TransaksiData{Nama: "Ayam Bakar Spesial Lebaran", Harga: 105000},
		},
	}
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
