package main

import "fmt"

type BangunDatar interface {
	Luas() float64
	Keliling() float64
}

type InputUntukPersegi struct {
	Sisi int
}

type InputUntukPersegiPanjang struct {
	Panjang int
	Lebar   int
}

// function untuk hitung persegi
func (input InputUntukPersegi) Luas() float64 {
	var luas = input.Sisi * input.Sisi
	return float64(luas)
}

func (input InputUntukPersegi) Keliling() float64 {
	keliling := input.Sisi * 4
	return float64(keliling)
}

// function untuk hitung persegi panjang
func (input InputUntukPersegiPanjang) Luas() float64 {
	luas := input.Panjang * input.Lebar
	return float64(luas)
}

func (input InputUntukPersegiPanjang) Keliling() float64 {
	keliling := 2*input.Panjang + 2*input.Lebar
	return float64(keliling)
}

func main() {
	dataPersegi := InputUntukPersegi{
		Sisi: 10,
	}

	dataPersegiPanjang := InputUntukPersegiPanjang{
		Panjang: 12,
		Lebar:   7,
	}

	luasPersegi := dataPersegi.Luas()
	luasPersegiPanjang := dataPersegiPanjang.Luas()
	kelilingPersegi := dataPersegi.Keliling()
	kelilingPersegiPanjang := dataPersegiPanjang.Keliling()
	fmt.Println(luasPersegi)
	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegi)
	fmt.Println(kelilingPersegiPanjang)

}
