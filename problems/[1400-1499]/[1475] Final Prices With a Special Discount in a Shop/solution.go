package main

func finalPrices(prices []int) []int {
	result := make([]int, len(prices))
	stack := make([]int, 0, len(prices))

	for i := len(prices) - 1; i >= 0; i-- {
		price := prices[i]

		for len(stack) > 0 && stack[len(stack)-1] > price {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = price - stack[len(stack)-1]
		} else {
			result[i] = price
		}

		stack = append(stack, price)
	}

	return result
}
