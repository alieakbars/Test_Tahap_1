package main

import (
	"fmt"
	"math"
	"time"
)

func validasiCuti(jumlahCutiBersama int, tanggalJoinStr string, tanggalRencanaCutiStr string, durasiCuti int) (bool, string) {
	if durasiCuti > 3 {
		return false, "Cuti pribadi maksimal 3 hari berturutan"
	}

	layout := "2006-01-02"

	tanggalJoin, err1 := time.Parse(layout, tanggalJoinStr)
	tanggalRencanaCuti, err2 := time.Parse(layout, tanggalRencanaCutiStr)

	if err1 != nil || err2 != nil {
		return false, "Format tanggal salah. Gunakan YYYY-MM-DD"
	}

	totalCutiPribadi := 14 - jumlahCutiBersama

	tanggalBolehCuti := tanggalJoin.AddDate(0, 0, 180)

	if tanggalRencanaCuti.Before(tanggalBolehCuti) {
		return false, "Belum memenuhi syarat 180 hari masa kerja"
	}

	if tanggalRencanaCuti.Year() == tanggalJoin.Year() {
		akhirTahun := time.Date(tanggalJoin.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
		
		selisihJam := akhirTahun.Sub(tanggalBolehCuti).Hours()
		jumlahHari := selisihJam / 24

		kuotaProrata := math.Floor((jumlahHari / 365.0) * float64(totalCutiPribadi))

		if float64(durasiCuti) > kuotaProrata {
			return false, fmt.Sprintf("Kuota cuti tidak mencukupi. Sisa kuota Anda hanya %d hari", int(kuotaProrata))
		}
	} else {
		if durasiCuti > totalCutiPribadi {
			return false, fmt.Sprintf("Kuota cuti tidak mencukupi. Kuota tahunan Anda %d hari", totalCutiPribadi)
		}
	}

	return true, "Silakan ambil cuti, pengajuan disetujui"
}

func main() {
	fmt.Println("Case 1")
	status, alasan := validasiCuti(7, "2021-05-01", "2021-11-05", 1)
	fmt.Printf("Status: %t\nAlasan: %s\n\n", status, alasan)

	fmt.Println("Case 2")
	status, alasan = validasiCuti(7, "2021-05-01", "2021-10-27", 1)
	fmt.Printf("Status: %t\nAlasan: %s\n\n", status, alasan)

	fmt.Println("Case 3")
	status, alasan = validasiCuti(7, "2021-05-01", "2021-11-10", 2)
	fmt.Printf("Status: %t\nAlasan: %s\n\n", status, alasan)

	fmt.Println("Case 4")
	status, alasan = validasiCuti(7, "2021-05-01", "2022-02-10", 4) 
	fmt.Printf("Status: %t\nAlasan: %s\n\n", status, alasan)
}