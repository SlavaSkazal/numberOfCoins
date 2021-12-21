package main

import (
	"fmt"
)

func main() {

	var N, K, M int

	fmt.Scanf("%d %d %d", &N, &K, &M)

	countingCoins(N, K, M)
}

func countingCoins(totalWeight, workpieceWeight, coinWeight int) {

	numberOfСoins := 0

	if totalWeight > 0 && totalWeight >= workpieceWeight && workpieceWeight >= coinWeight {
		numberOfСoins = weightCalculation(totalWeight, workpieceWeight, coinWeight)
	}

	fmt.Println(numberOfСoins)
}

func weightCalculation(totalWeight, workpieceWeight, coinWeight int) int {

	numberOfСoins := workpieceWeight / coinWeight

	remainder := workpieceWeight%coinWeight + (totalWeight - workpieceWeight)

	if remainder >= workpieceWeight {
		numberOfСoins += weightCalculation(remainder, workpieceWeight, coinWeight)
	}

	return numberOfСoins
}
