package main

import (
	"sort"
	"strconv"
	"strings"
)

func containChar(str1, str2 string) bool {
	var (
		tmpArr1 []string
		tmpArr2 []string
	)

	for _, val := range str1 {
		s := strconv.QuoteRuneToASCII(val)
		s = strings.Trim(s, "'")
		tmpArr1 = append(tmpArr1, s)
	}

	for _, val := range str2 {
		s := strconv.QuoteRuneToASCII(val)
		s = strings.Trim(s, "'")
		tmpArr2 = append(tmpArr2, s)
	}

	sort.Strings(tmpArr1)
	sort.Strings(tmpArr2)

	str1 = strings.Join(tmpArr1, "")
	str2 = strings.Join(tmpArr2, "")

	return str1 == str2
}

func getAnagram(dict []string) (tempIn []string) {
	for i := 0; i < len(dict); i++ {
		for j, val := range dict {
			if dict[i] == val {
				continue
			} else if len(dict[i]) != len(val) {
				continue
			} else {
				if containChar(dict[i], val) {
					if len(tempIn) == 0 {
						tempIn = append(tempIn, dict[i])
						tempIn = append(tempIn, val)
					} else {
						tempIn = append(tempIn, val)
					}
				}

				if j == (len(dict)-1) && len(tempIn) != 0 {
					return tempIn
				}
			}
		}
	}

	return tempIn
}
