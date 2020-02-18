package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tmicheletto/profiteer/lib"
)

func createIntSlice(nums []string) ([]int, error) {
	var r []int
	for _, v := range nums {
		i, err := strconv.Atoi(v)
		if err != nil {
			return r, err
		}
		r = append(r, i)
	}
	return r, nil
}

func main() {
	fmt.Print("Enter stock prices in AUD in a comma delimited list without spaces. e.g 1,2,3,4,5: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), ",")
	prices, err := createIntSlice(parts)
	if err != nil {
		fmt.Printf("Oops something went wrong parsing the stock prices you entered. %v\n", err)
		return
	}

	profit, err := lib.GetMaxProfit_SingleTransaction(prices)
	if err != nil {
		fmt.Printf("Oops something went wrong. %v\n", err)
		return
	}
	if profit > 0 {
		fmt.Printf("The maximum profit for those stock prices is $%d AUD.\n", profit)
	} else {
		fmt.Print("No profit could be made for those stock prices.\n")
	}
}
