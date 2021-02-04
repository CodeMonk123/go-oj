package main

import (
	"fmt"
)



func parseUnsignedInteger(seq []byte) bool {
	fmt.Printf("Try to parse unsigned from %s\n", string(seq))
	for _, ch := range seq {
		if ch >= '0' && ch <= '9' {
			continue
		}
		return false
	}
	return true
}

func parseInteger(seq []byte) bool {
	fmt.Printf("Try to parse integer from %s\n", string(seq))
	if seq[0] == '+' || seq[0] == '-' {
		if len(seq) == 1 {
			return false
		}
		return parseUnsignedInteger(seq[1:])
	}
	return parseUnsignedInteger(seq)
}

func parseUnsignedDecimal(seq []byte) bool {
	fmt.Printf("Try to parse unsigned decimal from %s\n", string(seq))
	for i, ch := range seq {
		if ch >= '0' && ch <= '9' {
			continue
		}
		if ch == '.' {
			if i == 0 {
				if len(seq) > 1 {
					return parseUnsignedInteger(seq[i+1:])
				} else {
					return false
				}
			}

			if i == len(seq) - 1 {
				if len(seq) > 1 {
					return true
				} else {
					return false
				}
			}
			return parseUnsignedInteger(seq[i+1:])
		}
		return false
	}
	return true
}

func parseDecimal(seq []byte) bool {
	fmt.Printf("Try to parse decimal from %s\n", string(seq))
	if len(seq) == 0 {
		return false
	}

	if seq[0] == '+' || seq[0] == '-' {
		if len(seq) == 1 {
			return false
		}
		return parseUnsignedDecimal(seq[1:])
	}
	return parseUnsignedDecimal(seq)
}

func parse(seq []byte) bool {
	// fmt.Printf("Try to parse from %s\n", string(seq))
	index := -1
	for i:=0;i < len(seq); i++ {
		if seq[i] == 'e' {
			if index == -1 {
				index = i
			} else {
				return false
			}
		}
	}
	if index == -1 {
		return parseDecimal(seq)
	}
	if index == 0 || index == len(seq) - 1 {
		return false
	}
	return parseDecimal(seq[0:index]) && parseInteger(seq[index+1:])
}

func isNumber(s string) bool {
	seq := []byte(s)
	start := 0
	end := len(seq) -1
	for start < len(seq) {
		if seq[start] == ' ' {
			start++
		} else {
			break
		}
	}

	for end >= 0 {
		if seq[end] == ' ' {
			end--
		} else {
			break
		}
	}

	if start > end {
		return false
	}

	// fmt.Println(string(seq[start:end+1]))

	return parse(seq[start:end+1])
}


func main() {
	fmt.Println(isNumber("0"))
	fmt.Println(isNumber(" 0.1"))
	fmt.Println(isNumber("abc"))
	fmt.Println(isNumber("1 a"))
	fmt.Println(isNumber("2e10"))
	fmt.Println(isNumber("-90e3"))
	fmt.Println(isNumber(" 1e"))
	fmt.Println(isNumber("e3"))
	fmt.Println(isNumber("6e-1"))
	fmt.Println(isNumber("99e2.5"))
	fmt.Println(isNumber("53.5e93"))
	fmt.Println(isNumber("3."))
	fmt.Println(isNumber("1e."))
	fmt.Println(isNumber("-e58"))
}
