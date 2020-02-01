// Given a string S of '(' and ')' parentheses, we add the minimum number of parentheses ( '(' or ')', and in any positions ) so that the resulting parentheses string is valid.

// Formally, a parentheses string is valid if and only if:

// It is the empty string, or
// It can be written as AB (A concatenated with B), where A and B are valid strings, or
// It can be written as (A), where A is a valid string.
// Given a parentheses string, return the minimum number of parentheses we must add to make the resulting string valid.

// Example 1:

// Input: "())"
// Output: 1
// Example 2:

// Input: "((("
// Output: 3
// Example 3:

// Input: "()"
// Output: 0
// Example 4:

// Input: "()))(("
// Output: 4

// Note:

// S.length <= 1000
// S only consists of '(' and ')' characters.

package main

import "fmt"

type stack struct {
	content []byte
}

func (s *stack) push(ch byte) {
	s.content = append(s.content, ch)
}

func (s *stack) top() byte {
	if len(s.content) == 0 {
		return ' '
	}
	return s.content[len(s.content)-1]
}

func (s *stack) pop() {
	s.content = s.content[:len(s.content)-1]
}

func (s *stack) length() int {
	return len(s.content)
}

func (s *stack) output() {
	for _, ch := range s.content {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}

func minAddToMakeValid(S string) int {
	s := new(stack)
	contents := []byte(S)
	for _, ch := range contents {
		if ch == '(' || s.top() == ' ' || s.top() == ')' {
			s.push(ch)
		} else {
			s.pop()
		}
		//s.output()
	}
	return s.length()
}

func main() {
	fmt.Println(minAddToMakeValid("()))(("))
}
