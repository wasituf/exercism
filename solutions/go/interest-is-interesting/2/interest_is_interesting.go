package interest

func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance < 1000:
		return 0.5
	case balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) * balance / 100
}

func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int

	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
	}

	return years
}
