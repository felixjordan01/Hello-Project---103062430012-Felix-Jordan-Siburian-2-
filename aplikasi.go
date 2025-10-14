package main

import (
	"fmt"
	"sort"
)

type Aplikasi struct {
	Nama    string
	Rating  float64
	Votes   int
}

func main() {
	var n int
	fmt.Print("Masukkan banyaknya aplikasi: ")
	fmt.Scanln(&n)

	aplikasi := make([]Aplikasi, 0)

	for i := 0; i < n; i++ {
		var nama string
		var rating float64
		var votes int
		fmt.Printf("Masukkan nama aplikasi %d: ", i+1)
		fmt.Scanln(&nama)
		fmt.Printf("Masukkan rating aplikasi %d: ", i+1)
		fmt.Scanln(&rating)
		fmt.Printf("Masukkan jumlah votes aplikasi %d: ", i+1)
		fmt.Scanln(&votes)

		if votes >= 10e6 {
			aplikasi = append(aplikasi, Aplikasi{Nama: nama, Rating: rating, Votes: votes})
		}
	}

	if len(aplikasi) == 0 {
		fmt.Println("Tidak ada aplikasi yang memenuhi syarat.")
		return
	}

	fmt.Println("Data Aplikasi:")
	for _, app := range aplikasi {
		fmt.Printf("Nama: %s, Rating: %.2f, Votes: %d\n", app.Nama, app.Rating, app.Votes)
	}

	rataRataRating := hitungRataRataRating(aplikasi)
	fmt.Printf("Rata-rata rating: %.2f\n", rataRataRating)

	namaAplikasiTertinggi := cariNamaAplikasiTertinggi(aplikasi)
	fmt.Printf("Nama aplikasi dengan votes tertinggi: %s\n", namaAplikasiTertinggi)
}

func hitungRataRataRating(aplikasi []Aplikasi) float64 {
	var totalRating float64
	for _, app := range aplikasi {
		totalRating += app.Rating
	}
	return totalRating / float64(len(aplikasi))
}

func cariNamaAplikasiTertinggi(aplikasi []Aplikasi) string {
	sort.Slice(aplikasi, func(i, j int) bool {
		return aplikasi[i].Votes > aplikasi[j].Votes
	})
	return aplikasi[0].Nama
}
