package Bangun3D

import "math"

type VolumeTabung struct {
	Phi    float64
	R      int
	T      int
	Volume float64
}

func HitungVolumeTabung(r int, t int) VolumeTabung {
	vol := math.Pi * math.Pow(float64(r), 2) * float64(t)
	volumeStruct := VolumeTabung{
		Phi:    math.Pi,
		R:      r,
		T:      t,
		Volume: vol,
	}

	return volumeStruct
}
