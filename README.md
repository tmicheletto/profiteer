# Profiteer

Solution to stock prices excerise described [here](https://gist.github.com/jonog/54e46b5b1200758d222e3c4cf61baaa6)

# Dependencies

GO 1.13

# To Run

In the root of the repo run `go install`. Then enter `profiteer`. You will be prompted to enter the stock prices in the following format `1,2,3,4,5`.

# To Test

Execute `go test ./...`.

To add tests, add an entry to the [test](./lib/stocks_test.go) table in the following format
```
{[]int{1, 2, 3, 4, 5}, 4, nil},
```
where the first item is the arrray of stock prices, the second item is the expected profit and the third is the expected error.