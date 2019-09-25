package main

func highestProfit(profits []int) (value int) {
	j := 1
	lowNumber := profits[0]
	for i := 0; i < len(profits); i++ {
		if j != len(profits) {
			if profits[i] < profits[j] {
				if lowNumber == profits[i] {
					lowNumber = profits[i]
				}

				if value < (profits[j] - lowNumber) {
					value = profits[j] - lowNumber
				}
			} else {
				lowNumber = profits[j]
			}
		}
		j++
	}

	return value
}
