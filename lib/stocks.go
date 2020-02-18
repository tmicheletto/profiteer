package lib

import "errors"

func GetMaxProfit_SingleTransaction(prices []int) (int, error) {
	if len(prices) < 2 {
		return -1, errors.New("there must be at least 2 prices")
	}
	maxProfit := -1
	minBuyPrice := -1

	for i, price := range prices[:len(prices)-1] {
		if price < 0 {
			return -1, errors.New("stock prices must be positive integers")
		}

		sellPrice := prices[i+1]

		if sellPrice-price > 0 {
			// Couldn't seem to figure out a better way to check if the minBuyPrice had been set.
			if minBuyPrice == -1 {
				minBuyPrice = price
			}

			if price < minBuyPrice {
				minBuyPrice = price
			}

			currentProfit := sellPrice - minBuyPrice
			// Couldn't seem to figure out a better way to check if the maxProfit had been set.
			if maxProfit == -1 {
				maxProfit = currentProfit
			}

			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}
	return maxProfit, nil
}

func GetMaxProfit_MultipleTransactions(prices []int) (int, error) {
	if len(prices) < 2 {
		return -1, errors.New("there must be at least 2 prices")
	}
	maxProfit := 0

	for i, price := range prices[:len(prices)-1] {
		if price < 0 {
			return -1, errors.New("stock prices must be positive integers")
		}

		sellPrice := prices[i+1]

		if sellPrice-price > 0 {
			currentProfit := sellPrice - price

			maxProfit = maxProfit + currentProfit
		}
	}
	return maxProfit, nil
}
