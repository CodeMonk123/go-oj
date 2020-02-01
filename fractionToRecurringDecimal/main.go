package main

import "strconv"

import "fmt"

import "math"

func contains(s []int, e int) (bool, int) {
	for i, v := range s {
		if e == v {
			return true, i
		}
	}
	return false, -1
}

func generateIntegerPart(numerator int, denominator int) string {
	return strconv.FormatInt(int64(numerator/denominator), 10)
}

func generateDecimalPart(numerator int, denominator int) string {
	numerator = numerator % denominator
	usedNumerator := append([]int{}, numerator*10)
	if numerator == 0 {
		return ""
	}
	next := numerator * 10 / denominator
	numerator = ((numerator * 10) % denominator) * 10
	res := append([]byte{}, '.')
	for numerator != 0 {
		res = append(res, byte('0'+next))
		if find, index := contains(usedNumerator, numerator); find {
			return "." + string(res[1:index+1]) + "(" + string(res[index+1:]) + ")"
		}
		usedNumerator = append(usedNumerator, numerator)
		next = numerator / denominator
		numerator = (numerator % denominator) * 10
	}
	res = append(res, byte('0'+next))
	return string(res)
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator*denominator >= 0 {
		return generateIntegerPart(int(math.Abs(float64(numerator))), int(math.Abs(float64(denominator)))) + generateDecimalPart(int(math.Abs(float64(numerator))), int(math.Abs(float64(denominator))))
	} else {
		return "-" + generateIntegerPart(int(math.Abs(float64(numerator))), int(math.Abs(float64(denominator)))) + generateDecimalPart(int(math.Abs(float64(numerator))), int(math.Abs(float64(denominator))))
	}

}

func main() {
	fmt.Println(fractionToDecimal(1, 6))
	fmt.Println(fractionToDecimal(1, 7))
	fmt.Println(fractionToDecimal(1, 3))
	fmt.Println(fractionToDecimal(3, 5))
	fmt.Println(fractionToDecimal(-50, 8))
}
