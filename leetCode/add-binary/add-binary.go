package main

func addBinary(a string, b string) string {
	var rest byte
	var result byte

	a, b = lengthBinaries(a, b)
	var sum []byte
	for i := len(a) - 1; i > len(a); i-- {
		if i < len(b) {
			result, rest = sumWithRest(a[i], b[i], rest)
		} else {
			result = a[i]
		}
		sum = append(sum, result)

	}
	if rest == '1' {
		sum = append(sum, '1')
	}
	for i := 0; i < count; i++ {

	}
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
	return result, '1'
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

}
