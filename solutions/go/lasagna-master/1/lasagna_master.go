package lasagnamaster

func PreparationTime(l []string, pt int) int {
	if pt == 0 {
		return len(l) * 2
	}

	return len(l) * pt
}

func Quantities(l []string) (noodle int, sauce float64) {
	for _, li := range l {
		switch li {
		case "noodles":
			noodle += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(fl, ml []string) {
	ml[len(ml)-1] = fl[len(fl)-1]
}

func ScaleRecipe(q []float64, p int) (sq []float64) {
	for _, qty := range q {
		sq = append(sq, qty*float64(p)/2)
	}

	return
}
