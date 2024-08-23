package shunting_yard

import (
	"fmt"

	"dsa-exploration/stack"
)

var operatorPrecedenceMap = map[string]int{
	"+": 2,
	"-": 2,
	"x": 3,
	"/": 3,
	"^": 4,
}

type operatorAssociativity int

const (
	operatorAssociativityLeft = iota
	operatorAssociativityRight
)

var operatorAssociativityMap = map[string]operatorAssociativity{
	"+": operatorAssociativityLeft,
	"-": operatorAssociativityLeft,
	"x": operatorAssociativityLeft,
	"/": operatorAssociativityLeft,
	"^": operatorAssociativityRight,
}

func isNumber(token rune) bool {
	return token >= 48 && token <= 57
}

func isLeftParentheses(token string) bool {
	return token == "("
}

func isRightParentheses(token string) bool {
	return token == ")"
}

func getPostfixFormula(formula string) (postfixFormula []string) {
	operatorStack := stack.CreateStack()

	for _, token := range formula {
		tokenStr := fmt.Sprintf("%s", string(token))

		if isNumber(token) {
			postfixFormula = append(postfixFormula, tokenStr)
		} else {
			if operatorStack.IsEmpty() {
				operatorStack.Push(tokenStr)
				continue
			}

			if !isRightParentheses(tokenStr) {
				for !operatorStack.IsEmpty() {
					topStack := operatorStack.Pop()

					if !isLeftParentheses(tokenStr) &&
						operatorPrecedenceMap[topStack.(string)] >= operatorPrecedenceMap[tokenStr] &&
						operatorAssociativityMap[tokenStr] == operatorAssociativityLeft {

						postfixFormula = append(postfixFormula, topStack.(string))
					} else {
						operatorStack.Push(topStack)
						break
					}
				}
				operatorStack.Push(tokenStr)
			} else {
				for !operatorStack.IsEmpty() {
					topStack := operatorStack.Pop()

					if isLeftParentheses(topStack.(string)) {
						break
					}

					postfixFormula = append(postfixFormula, topStack.(string))
				}
			}
		}
	}

	for !operatorStack.IsEmpty() {
		topStack := operatorStack.Pop()

		postfixFormula = append(postfixFormula, topStack.(string))
	}

	return postfixFormula
}
