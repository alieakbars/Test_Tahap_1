package main

import (
	"fmt"
)

func hitungKembalian(totalBelanja, uangDibayar int) {
	fmt.Printf("Total belanja : Rp %d\n", totalBelanja)
	fmt.Printf("Pembeli membayar: Rp %d\n\n", uangDibayar)
	fmt.Println("Output:")

	if uangDibayar < totalBelanja {
		fmt.Println("False, kurang bayar\n")
		return
	}

	kembalianAsli := uangDibayar - totalBelanja
	kembalianDibulatkan := (kembalianAsli / 100) * 100

	fmt.Printf("Kembalian yang harus diberikan kasir: %d, dibulatkan menjadi %d\n\n", kembalianAsli, kembalianDibulatkan)

	if kembalianDibulatkan == 0 {
		return
	}

	fmt.Println("Pecahan uang:")
	pecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	sisaKembalian := kembalianDibulatkan
	for _, p := range pecahan {
		jumlah := sisaKembalian / p

		if jumlah > 0 {
			satuan := "lembar"
			if p < 1000 {
				satuan = "koin"
			}
			fmt.Printf("%d %s %d\n", jumlah, satuan, p)
			
			sisaKembalian = sisaKembalian % p
		}
	}
}

func main() {
	fmt.Println("Input 1:")
	hitungKembalian(700649, 800000)

	fmt.Println("Input 2:")
	hitungKembalian(575650, 580000)

	fmt.Println("Input 3:")
	hitungKembalian(657650, 600000)
}