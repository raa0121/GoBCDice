package notation

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/ast"
)

// nodeの中置記法表記を返す
func InfixNotation(node ast.Node) (string, error) {
	switch n := node.(type) {
	case *ast.Calc:
		return infixNotationOfCalc(n)
	case ast.PrefixExpression:
		return infixNotationOfPrefixExpression(n)
	case ast.InfixExpression:
		return infixNotationOfInfixExpression(n)
	case *ast.Int:
		return fmt.Sprintf("%d", n.Value), nil
	}

	return "", fmt.Errorf("infix notation not implemented: %s", node.Type())
}

func infixNotationOfCalc(node *ast.Calc) (string, error) {
	expr, err := InfixNotation(node.Expression())
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("C(%s)", expr), nil
}

func infixNotationOfPrefixExpression(node ast.PrefixExpression) (string, error) {
	right := node.Right()
	rightInfixNotation, rightErr := InfixNotation(right)
	if rightErr != nil {
		return "", rightErr
	}

	if right.Type() == ast.INT_NODE {
		return node.Operator() + rightInfixNotation, nil
	}

	return fmt.Sprintf("%s(%s)", node.Operator(), rightInfixNotation), nil
}

func infixNotationOfInfixExpression(node ast.InfixExpression) (string, error) {
	left := node.Left()
	leftInfixNotation, leftErr := InfixNotation(left)
	if leftErr != nil {
		return "", leftErr
	}

	right := node.Right()
	rightInfixNotation, rightErr := InfixNotation(right)
	if rightErr != nil {
		return "", rightErr
	}

	var parenthesizedLeft string
	switch l := left.(type) {
	case ast.InfixExpression:
		if l.Precedence() < node.Precedence() {
			parenthesizedLeft = fmt.Sprintf("(%s)", leftInfixNotation)
		} else {
			parenthesizedLeft = leftInfixNotation
		}
	default:
		parenthesizedLeft = leftInfixNotation
	}

	var parenthesizedRight string
	switch r := right.(type) {
	case ast.PrefixExpression:
		parenthesizedRight = fmt.Sprintf("(%s)", rightInfixNotation)
	case ast.InfixExpression:
		lowPrecedence := r.Precedence() < node.Precedence()
		samePrecedenceAndNonCommutative :=
			r.Precedence() == node.Precedence() && !node.IsCommutative()

		if lowPrecedence || samePrecedenceAndNonCommutative {
			parenthesizedRight = fmt.Sprintf("(%s)", rightInfixNotation)
		} else {
			parenthesizedRight = rightInfixNotation
		}
	default:
		parenthesizedRight = rightInfixNotation
	}

	return parenthesizedLeft + node.Operator() + parenthesizedRight, nil
}
