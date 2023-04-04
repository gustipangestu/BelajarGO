package main

import (
	"fmt"
	"math"
)

type Bangun2D interface {
	Luas()
}

type InputanSegitiga struct {
	Alas   int
	Tinggi int
}

type InputanLingkaran struct {
	Radius int
}

// hitung luas segitiga
func (input InputanSegitiga) Luas() float64 {
	luas := 0.5 * float64(input.Alas) * float64(input.Tinggi)
	return luas
}

// hitung nilai luas lingkaran
func (input InputanLingkaran) Luas() float64 {
	luas := math.Pi * math.Pow(float64(input.Radius), 2)
	return luas
}

func main() {
	inputSegitiga := InputanSegitiga{
		Alas:   10,
		Tinggi: 8,
	}
	inputLingkaran := InputanLingkaran{
		Radius: 13,
	}
	luasSegitiga := inputSegitiga.Luas()
	luasLingkaran := inputLingkaran.Luas()
	fmt.Println(luasSegitiga)
	fmt.Println(luasLingkaran)
}
