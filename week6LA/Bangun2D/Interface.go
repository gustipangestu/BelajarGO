package Bangun2D

import "math"

type TestInterface interface {
	// Methods
	AreaTest() float64
	KelilingTest() float64
}

type Kotak struct {
	Sisi int
}

type Circle struct {
	Radius float64
}

func (k Kotak) AreaTest() float64 {
	luas := k.Sisi * k.Sisi
	return float64(luas)
}

func (c Circle) AreaTest() float64 {
	luas := math.Pi * math.Pow(c.Radius, 2)
	return luas
}

func (k Kotak) KelilingTest() float64 {
	kel := k.Sisi * 4
	return float64(kel)
}

// ==========================================================

func Interfacee(t TestInterface) float64 {
	luasKotak := t.AreaTest()

	return luasKotak
}
