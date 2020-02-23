package main

import "fmt"

//Given a array of numbers representing the stock prices of a company in chronological order, write a function that calculates the maximum profit you could have made from buying and selling that stock once. You must buy before you can sell it.
func main() {
	prices := []int{9, 11, 8, 5, 7, 10}
	fmt.Println(maxProfit(prices))
}

func maxProfit(price []int) int {
	profit := 0
	if len(price) == 0 {
		return profit
	}
	min, max := price[0], price[0]
	for i := range price {
		if price[i] > max {
			max = price[i]
		}

		if price[i] < min {
			min = price[i]
			max = price[i]
		}

		if max-min > profit {
			profit = max - min
		}
	}

	return profit
}
