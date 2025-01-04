package main

import (
	"fmt"
	"sort"

	"github.com/ephrain1021/go-modules/set"
	"github.com/expr-lang/expr"
)

func tryExpression(allResults map[string]*set.Set, a, b, c, d int, expression string) {
	result, err := expr.Eval(expression, nil)
	if err == nil {
		if value, ok := result.(int); ok && value == 10 {
			// Sort the digits with sort function
			digits := []int{a, b, c, d}
			sort.Ints(digits)
			key := fmt.Sprintf("%d%d%d%d", digits[0], digits[1], digits[2], digits[3])
			if _, ok := allResults[key]; !ok {
				allResults[key] = set.New()
			}
			allResults[key].Add(expression)
		}
	}
}

func genCombinationsWith0Parentheses(allResults map[string]*set.Set, a, b, c, d int, operators []string) {
	for _, op1 := range operators {
		for _, op2 := range operators {
			for _, op3 := range operators {
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("%d %s %d %s %d %s %d", a, op1, b, op2, c, op3, d))
			}
		}
	}
}

func genCombinationsWith1Parentheses(allResults map[string]*set.Set, a, b, c, d int, operators []string) {
	for _, op1 := range operators {
		for _, op2 := range operators {
			for _, op3 := range operators {
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("(%d %s %d) %s %d %s %d", a, op1, b, op2, c, op3, d))
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("%d %s (%d %s %d) %s %d", a, op1, b, op2, c, op3, d))
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("%d %s %d %s (%d %s %d)", a, op1, b, op2, c, op3, d))

				tryExpression(allResults, a, b, c, d, fmt.Sprintf("(%d %s %d %s %d) %s %d", a, op1, b, op2, c, op3, d))
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("%d %s (%d %s %d %s %d)", a, op1, b, op2, c, op3, d))
			}
		}
	}
}

func genCombinationsWith2Parentheses(allResults map[string]*set.Set, a, b, c, d int, operators []string) {
	for _, op1 := range operators {
		for _, op2 := range operators {
			for _, op3 := range operators {
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("(%d %s %d) %s (%d %s %d)", a, op1, b, op2, c, op3, d))

				tryExpression(allResults, a, b, c, d, fmt.Sprintf("((%d %s %d) %s %d) %s %d", a, op1, b, op2, c, op3, d))
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("(%d %s (%d %s %d)) %s %d", a, op1, b, op2, c, op3, d))
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("%d %s ((%d %s %d) %s %d)", a, op1, b, op2, c, op3, d))
				tryExpression(allResults, a, b, c, d, fmt.Sprintf("%d %s (%d %s (%d %s %d))", a, op1, b, op2, c, op3, d))
			}
		}
	}
}

func main() {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	operators := []string{"+", "-", "*", "/"}
	allResults := make(map[string]*set.Set)

	for _, a := range digits {
		for _, b := range digits {
			for _, c := range digits {
				for _, d := range digits {
					genCombinationsWith0Parentheses(allResults, a, b, c, d, operators)
					genCombinationsWith1Parentheses(allResults, a, b, c, d, operators)
					genCombinationsWith2Parentheses(allResults, a, b, c, d, operators)
				}
			}
		}
	}

	fmt.Println(len(allResults))

	// Print all results, each expression in one line
	for key, expression_list := range allResults {
		fmt.Println(key)
		for _, expression := range expression_list.List() {
			fmt.Printf("  %s\n", expression)
		}
	}
}
