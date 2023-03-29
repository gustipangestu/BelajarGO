package Bangun2D

type DataLuasPersegi struct {
	Sisi int
	Luas float64
}

func LuasPersegi(sisi int) DataLuasPersegi {
	var luas = sisi * sisi
	var dataLuas = DataLuasPersegi{
		Sisi: sisi,
		Luas: float64(luas),
	}
	return dataLuas
}
