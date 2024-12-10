package weightconv

func KToP(k Kilogram) Pound {
	return Pound(k * 2.20462)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p / 2.20462)
}
