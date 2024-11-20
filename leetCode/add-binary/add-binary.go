package main

import "fmt"

func addBinary(a string, b string) string {
	var rest byte
	var result byte

	rest = '0'
	a, b = lengthBinaries(a, b)
	var sum string
	j := len(b) - 1
	for i := len(a) - 1; i >= 0; i-- {
		if j >= 0 {
			result, rest = sumWithRest(a[i], b[j], rest)
			j--
		} else {
			result, rest = sumRules(a[i], rest)
		}
		sum = string(result) + sum

	}

	if rest == '1' {
		sum = string(rest) + sum
	}
	return sum
}

func lengthBinaries(a, b string) (string, string) {
	if len(a) > len(b) {
		return a, b
	}
	return b, a
}

func sumWithRest(a, b, c byte) (result, rest byte) {
	result, rest = sumRules(a, b)
	if c == '0' {
		return result, rest
	}
	if result == '0' {
		return '1', rest
	}
	return '0', '1'
}

func sumRules(a, b byte) (result, rest byte) {
	switch true {
	case a == '0' && b == '0':
		return '0', '0'
	case a == '1' && b == '0':
		return '1', '0'
	case a == '0' && b == '1':
		return '1', '0'
	case a == '1' && b == '1':
		return '0', '1'
	}
	return '0', '0'
}

func main() {

	a := "110010"
	b := "10111"

	/*
				 110010
				  10111
		right	1001001
				1001101

	*/
	sum := addBinary(a, b)
	fmt.Print(sum)
}
