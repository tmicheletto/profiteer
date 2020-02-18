package lib_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmicheletto/profiteer/lib"
)

func TestGetMaxProfit_SingleTransaction(t *testing.T) {
	tables := []struct {
		stocks []int
		profit int
		err    error
	}{
		{[]int{1}, -1, errors.New("there must be at least 2 prices")},
		{[]int{-1, 0, 1, 2, 3}, -1, errors.New("stock prices must be positive integers")},
		{[]int{1, 1}, -1, nil},
		{[]int{1, 2}, 1, nil},
		{[]int{10, 7, 5, 8, 11, 9}, 6, nil},
		{[]int{50, 50, 51, 51, 35, 30, 25, 20, 15, 10, 5, 6, 7, 8}, 3, nil},
		{[]int{50, 50, 50, 50, 50, 48, 45, 42, 39, 36, 33, 30, 20, 10, 5, 1}, -1, nil},
		{[]int{50, 51, 52, 50, 50, 48, 45, 42, 39, 36, 33, 30, 20, 10, 5, 1}, 2, nil},
		{[]int{5, 8, 11, 14, 11, 14, 17, 20, 17, 14, 11, 8, 5}, 15, nil},
	}

	for _, table := range tables {
		profit, err := lib.GetMaxProfit_SingleTransaction(table.stocks)

		assert.Equal(t, err, table.err)
		assert.Equal(t, profit, table.profit)
	}
}

func TestGetMaxProfit_MultipleTransactions(t *testing.T) {
	tables := []struct {
		stocks []int
		profit int
		err    error
	}{
		{[]int{1}, -1, errors.New("there must be at least 2 prices")},
		{[]int{-1, 0, 1, 2, 3}, -1, errors.New("stock prices must be positive integers")},
		{[]int{1, 1}, 0, nil},
		{[]int{1, 2}, 1, nil},
		{[]int{7, 5, 3, 6, 9, 4}, 6, nil},
	}

	for _, table := range tables {
		profit, err := lib.GetMaxProfit_MultipleTransactions(table.stocks)

		assert.Equal(t, err, table.err)
		assert.Equal(t, profit, table.profit)
	}
}
