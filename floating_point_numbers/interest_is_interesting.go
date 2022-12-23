package main

import "fmt"

var myBalance, myTargetBalance = 1000.0, 1032.682765146664

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var greedIsGood float32
	switch {
	case balance < 0:
		greedIsGood = 3.213
	case balance >= 0 && balance < 1000:
		greedIsGood = 0.5
	case balance >= 1000 && balance < 5000:
		greedIsGood = 1.621
	case balance >= 5000:
		greedIsGood = 2.475
	}
	return greedIsGood
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := float64(InterestRate(balance))
	rate = balance * (rate / 100)
	return rate
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	rate := Interest(balance)
	updatedBalance := balance + rate
	return updatedBalance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var yearCount int

	if balance == targetBalance {
		return 0
	}

	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		yearCount++
	}
	return yearCount
}

func main() {
	fmt.Println(InterestRate(myBalance))
	fmt.Println(Interest(myBalance))
	fmt.Println(AnnualBalanceUpdate(myBalance))
	fmt.Println(YearsBeforeDesiredBalance(myBalance, myTargetBalance))
}
