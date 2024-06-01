package world

func (w *World) Turn() {
	for _, country := range w.Countries {
		for i, human := range country.Human {
			if human <= 0 {
				country.Human[i] = 1
			} else {
				country.Human[i] = country.Human[i] * country.birthRate
			}
		}
	}
}
