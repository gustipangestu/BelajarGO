package Pria

type DataPria struct {
	Nama    string
	Tb      float64
	BbIdeal float64
}

func HitungBBIdealPria(nama string, tb float64) DataPria {
	bbIdeal := (tb - 100) - ((tb - 100) * 0.1)
	output := DataPria{
		Nama:    nama,
		Tb:      tb,
		BbIdeal: bbIdeal,
	}
	return output
}
