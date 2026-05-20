package main

import (
	"fmt"
	"strings"
)

func cariKecocokan(n int, data []string) string {
	targetFound := false
	var target string

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if strings.ToLower(data[i]) == strings.ToLower(data[j]) {
				target = strings.ToLower(data[i])
				targetFound = true
				break
			}
		}
		if targetFound {
			break
		}
	}

	if !targetFound {
		return "false"
	}

	var hasil string
	for i := 0; i < n; i++ {
		if strings.ToLower(data[i]) == target {
			if hasil == "" {
				hasil = fmt.Sprintf("%d", i+1)
			} else {
				hasil += fmt.Sprintf(" %d", i+1)
			}
		}
	}

	return hasil
}

func main() {
	contoh1 := []string{"abcd", "acbd", "aaab", "acbd"}
	fmt.Println("Contoh input 1:")
	fmt.Println("Output:", cariKecocokan(4, contoh1))

	contoh2 := []string{"pisang", "goreng", "enak", "sekali", "rasanya"}
	fmt.Println("Contoh input 2:")
	fmt.Println("Output:", cariKecocokan(5, contoh2))

	contoh3 := []string{
		"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate",
		"Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk",
	}
	fmt.Println("Contoh input 3:")
	fmt.Println("Output:", cariKecocokan(11, contoh3))
}
