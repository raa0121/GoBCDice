/*
抽象構文木の中置表記生成処理のパッケージ。

抽象構文木の中置表記は、BCDiceコマンドの実行結果のメッセージに含まれる。
中置表記を表示することで、利用者が入力したコマンドが構文解析器に正しく解釈されていることを示すことができる。
この中置表記は、日常的に数式で使われている、最小限の括弧を含むものとする。
*/
package notation

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"strings"
)

// InfixNotation は構文解析木の中置表記を返す。
func InfixNotation(node ast.Node) (string, error) {
	switch n := node.(type) {
	case *ast.DRollExpr:
		return infixNotationOfDRollExpr(n)
	case *ast.Calc:
		return infixNotationOfCalc(n)
	case ast.Divide:
		return infixNotationOfDivide(n)
	case ast.PrefixExpression:
		return infixNotationOfPrefixExpression(n)
	case ast.InfixExpression:
		return infixNotationOfInfixExpression(n)
	case *ast.Int:
		return fmt.Sprintf("%d", n.Value), nil
	case *ast.SumRollResult:
		return infixNotationOfSumRollResult(n)
	}

	return "", fmt.Errorf("infix notation not implemented: %s", node.Type())
}

// infixNotationOfCalc は加算ロール式の中置表記を返す。
func infixNotationOfDRollExpr(node *ast.DRollExpr) (string, error) {
	expr, err := InfixNotation(node.Expression())
	if err != nil {
		return "", err
	}

	return expr, nil
}

// infixNotationOfCalcは計算ノードの中置表記を返す
func infixNotationOfCalc(node *ast.Calc) (string, error) {
	expr, err := InfixNotation(node.Expression())
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("C(%s)", expr), nil
}

// infixNotationOfPrefixExpression は前置式の中置表記を返す。
func infixNotationOfPrefixExpression(node ast.PrefixExpression) (string, error) {
	right := node.Right()
	rightInfixNotation, rightErr := InfixNotation(right)
	if rightErr != nil {
		return "", rightErr
	}

	if right.IsPrimaryExpression() {
		return node.Operator() + rightInfixNotation, nil
	}

	return node.Operator() + Parenthesize(rightInfixNotation), nil
}

// parenthesizeChildOfInfixExpression は中置式の子ノードの中置表記を返す。
// 必要な場合は子ノードの中置表記を括弧で囲む。
//
// parent: 中置式のノード,
// child: 中置式の子ノード,
// parentIsAssociative: 中置式が子ノードの方向に結合性かどうか。
func parenthesizeChildOfInfixExpression(
	parent ast.InfixExpression,
	child ast.Node,
	parentIsAssociative bool,
) (string, error) {
	infixNotationOfChild, err := InfixNotation(child)
	if err != nil {
		return "", err
	}

	switch c := child.(type) {
	case ast.PrefixExpression:
		return Parenthesize(infixNotationOfChild), nil
	case ast.InfixExpression:
		lowPrecedence := c.Precedence() < parent.Precedence()
		samePrecedenceAndNonAssociative :=
			c.Precedence() == parent.Precedence() && !parentIsAssociative
		if lowPrecedence || samePrecedenceAndNonAssociative {
			return Parenthesize(infixNotationOfChild), nil
		}

		return infixNotationOfChild, nil
	default:
		return infixNotationOfChild, nil
	}
}

// Parenthesize は文字列を括弧で囲む。
func Parenthesize(s string) string {
	return "(" + s + ")"
}

// infixNotationOfInfixExpression は中置式の中置表記を返す。
func infixNotationOfInfixExpression(node ast.InfixExpression) (string, error) {
	left, right, err := infixNotationsOfInfixExpressionChildren(node)
	if err != nil {
		return "", err
	}

	return left + node.Operator() + right, nil
}

// infixNotationOfDivide は除算の中置表記を返す。
// 除算では端数処理の方法を除数の後で示す必要があるため、処理が特別になる。
func infixNotationOfDivide(node ast.Divide) (string, error) {
	left, right, err := infixNotationsOfInfixExpressionChildren(node)
	if err != nil {
		return "", err
	}

	return left + "/" + right + node.RoundingMethod().String(), nil
}

// infixNotationsOfInfixExpressionChildren は中置式の左右の子ノードの中置表記を返す。
func infixNotationsOfInfixExpressionChildren(node ast.InfixExpression) (string, string, error) {
	var leftInfixNotation string
	var leftErr error

	// 演算子が左結合性の場合、単項マイナスは特別扱い
	// 例えば、-1+2 の中置表記が (-1)+2 ではなく -1+2 とならなければならない
	//
	// 左結合性でない場合、例えば (-1)^2 の中置表記は (-1)^2 のままとなる
	if leftUMinus, uMinus := node.Left().(*ast.UnaryMinus); uMinus && node.IsLeftAssociative() {
		leftInfixNotation, leftErr = infixNotationOfPrefixExpression(leftUMinus)
	} else {
		leftInfixNotation, leftErr = parenthesizeChildOfInfixExpression(
			node,
			node.Left(),
			node.IsLeftAssociative(),
		)
	}
	if leftErr != nil {
		return "", "", leftErr
	}

	rightInfixNotation, rightErr := parenthesizeChildOfInfixExpression(
		node,
		node.Right(),
		node.IsRightAssociative(),
	)
	if rightErr != nil {
		return "", "", rightErr
	}

	return leftInfixNotation, rightInfixNotation, nil
}

// infixNotationOfSumRollResult は加算ロール結果の中置表記を返す。
func infixNotationOfSumRollResult(node *ast.SumRollResult) (string, error) {
	dieValueStrs := []string{}

	for _, d := range node.Dice {
		dieValueStrs = append(dieValueStrs, fmt.Sprintf("%d", d.Value))
	}

	return fmt.Sprintf("%d[%s]", node.Value(), strings.Join(dieValueStrs, ",")), nil
}
