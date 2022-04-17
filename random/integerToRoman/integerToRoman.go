package main

import "log"

/**
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

given 1 <= num <= 3999
*/

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i, v := range vals {
		for v <= num {
			res += symbols[i]
			num -= v
		}
	}
	return res
}

func intToRoman2(num int) string {
	var res string
	for num > 0 {
		if t := num / 1000; t > 0 {
			res = addNCharAfter(res, t, "M")
			num = num % 1000
			continue
		}

		if t := num / 100; t > 0 {
			res += convertIntToRoman(t, "M", "D", "C")
			num = num % 100
			continue
		}

		if t := num / 10; t > 0 {
			res += convertIntToRoman(t, "C", "L", "X")
			num = num % 10
			continue
		}
		if num < 10 {
			res += convertIntToRoman(num, "X", "V", "I")
			num /= 10
		}
	}
	return res
}

func convertIntToRoman(numToConvert int, charForHighest, charForMid, charForUnit string) string {
	var res string
	if numToConvert >= 1 && numToConvert <= 3 {
		res = addNCharAfter("", numToConvert, charForUnit)
	}
	if numToConvert == 4 {
		res = charForUnit + charForMid
	}
	if numToConvert == 5 {
		res = charForMid
	}
	if numToConvert >= 6 && numToConvert <= 8 {
		res = addNCharAfter(charForMid, numToConvert-5, charForUnit)
	}
	if numToConvert == 9 {
		res = charForUnit + charForHighest
	}
	return res

}

func addNCharAfter(origin string, n int, char string) string {
	for i := 0; i < n; i++ {
		origin += char
	}
	return origin
}

func main() {
	log.Println(intToRoman(1994))
	log.Println(intToRoman(345))
	log.Println(intToRoman(58))
	log.Println(intToRoman(7))

}
