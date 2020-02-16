package profiteer

import "errors"

//GetMaxProfit returns the maximum profit possible to be made
//from 1 purchase and 1 sale of a given set of stock prices
func GetMaxProfit(prices []int) (int, error) {
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
			if minBuyPrice == -1 {
				minBuyPrice = price
			}

			if price < minBuyPrice {
				minBuyPrice = price
			}

			currentProfit := sellPrice - minBuyPrice
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
