package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dictSoal3 := []string{"bat", "code", "cat", "act", "cab", "crazy", "tac"}
	soalNo3 := getAnagram(dictSoal3)
	fmt.Printf("Result Soal Nomor 3 => %s\n", strings.Join(soalNo3, " "))

	dictSoal4 := []int{5, 6, 15, 3, 10, 22, 15}
	// dictSoal4 := []int{10, 15, 8, 7, 11}
	// dictSoal4 := []int{100, 90, 80, 75, 65}
	profit := highestProfit(dictSoal4)
	soalNo4 := strconv.Itoa(profit)
	if profit == 0 {
		soalNo4 = "<tidak dapat membeli>"
	}

	fmt.Printf("Result Soal Nomor 4 => %s\n", soalNo4)
}
