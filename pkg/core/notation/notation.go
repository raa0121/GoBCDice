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
//
// node: 構文解析木のルートノード,
// walkingToLeft: 左側への探索を続けているか。
//
// walkingToLeft は、単項マイナスを括弧で囲むかどうかの決定に使われる。
// 括弧で囲まれた範囲のルートノードから連続して左側を調べていき、
// 最左端が単項マイナスであれば括弧で囲まないように、そうでなければ
// 括弧で囲むようにしている。
// この最左端かどうかの判定にwalkingToLeftを使っている。
// ルートノードに対してこの関数を呼び出すときはwalkingToLeftをtrueに
// 設定し、右側の中置表記を求める際にはfalseを設定する。
func InfixNotation(node ast.Node, walkingToLeft bool) (string, error) {
	switch n := node.(type) {
	case *ast.Calc:
		return infixNotationOfCalc(n, walkingToLeft)
	case ast.Command:
		return infixNotationOfCommand(n, walkingToLeft)
	case *ast.Compare:
		return infixNotationOfCompare(n, walkingToLeft)
	case ast.Divide:
		return infixNotationOfDivide(n, walkingToLeft)
	case ast.PrefixExpression:
		return infixNotationOfPrefixExpression(n, walkingToLeft)
	case ast.InfixExpression:
		return infixNotationOfInfixExpression(n, walkingToLeft)
	case *ast.Int:
		return fmt.Sprintf("%d", n.Value), nil
	case *ast.SumRollResult:
		return infixNotationOfSumRollResult(n, walkingToLeft)
	}

	return "", fmt.Errorf("infix notation not implemented: %s", node.Type())
}

// infixNotationOfCommand はコマンドの中置表記を返す。
func infixNotationOfCommand(node ast.Command, walkingToLeft bool) (string, error) {
	expr, err := InfixNotation(node.Expression(), walkingToLeft)
	if err != nil {
		return "", err
	}

	return expr, nil
}

// infixNotationOfCalc は計算ノードの中置表記を返す。
func infixNotationOfCalc(node *ast.Calc, walkingToLeft bool) (string, error) {
	expr, err := InfixNotation(node.Expression(), walkingToLeft)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("C(%s)", expr), nil
}

// infixNotationOfCompare は比較式の中置表記を返す。
func infixNotationOfCompare(node *ast.Compare, _ bool) (string, error) {
	leftInfixNotation, leftErr := InfixNotation(node.Left(), true)
	if leftErr != nil {
		return "", leftErr
	}

	rightInfixNotation, rightErr := InfixNotation(node.Right(), true)
	if rightErr != nil {
		return "", rightErr
	}

	return leftInfixNotation + node.Operator() + rightInfixNotation, nil
}

// infixNotationOfPrefixExpression は前置式の中置表記を返す。
func infixNotationOfPrefixExpression(node ast.PrefixExpression, walkingToLeft bool) (string, error) {
	right := node.Right()

	if right.IsPrimaryExpression() {
		// 一次式の場合は括弧で囲まない
		rightInfixNotation, err := InfixNotation(right, walkingToLeft)
		if err != nil {
			return "", err
		}

		return node.Operator() + rightInfixNotation, nil
	}

	// 一次式でない場合は括弧で囲む
	rightInfixNotation, err := InfixNotation(right, true)
	if err != nil {
		return "", err
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
	walkingToLeft bool,
) (string, error) {
	switch c := child.(type) {
	case ast.PrefixExpression:
		{
			infixNotationOfChild, err := InfixNotation(child, walkingToLeft)
			if err != nil {
				return "", err
			}

			return Parenthesize(infixNotationOfChild), nil
		}
	case ast.InfixExpression:
		{
			lowPrecedence := c.Precedence() < parent.Precedence()
			samePrecedenceAndNonAssociative :=
				c.Precedence() == parent.Precedence() && !parentIsAssociative
			if lowPrecedence || samePrecedenceAndNonAssociative {
				infixNotationOfChild, err := InfixNotation(child, true)
				if err != nil {
					return "", err
				}
				return Parenthesize(infixNotationOfChild), nil
			}

			infixNotationOfChild, err := InfixNotation(child, walkingToLeft)
			if err != nil {
				return "", err
			}

			return infixNotationOfChild, nil
		}
	default:
		{
			infixNotationOfChild, err := InfixNotation(child, walkingToLeft)
			if err != nil {
				return "", err
			}

			return infixNotationOfChild, nil
		}
	}
}

// Parenthesize は文字列を括弧で囲む。
func Parenthesize(s string) string {
	return "(" + s + ")"
}

// infixNotationOfInfixExpression は中置式の中置表記を返す。
func infixNotationOfInfixExpression(node ast.InfixExpression, walkingToLeft bool) (string, error) {
	left, right, err := infixNotationsOfInfixExpressionChildren(node, walkingToLeft)
	if err != nil {
		return "", err
	}

	return left + node.Operator() + right, nil
}

// infixNotationOfDivide は除算の中置表記を返す。
// 除算では端数処理の方法を除数の後で示す必要があるため、処理が特別になる。
func infixNotationOfDivide(node ast.Divide, walkingToLeft bool) (string, error) {
	left, right, err := infixNotationsOfInfixExpressionChildren(node, walkingToLeft)
	if err != nil {
		return "", err
	}

	return left + "/" + right + node.RoundingMethod().String(), nil
}

// infixNotationsOfInfixExpressionChildren は中置式の左右の子ノードの中置表記を返す。
func infixNotationsOfInfixExpressionChildren(node ast.InfixExpression, walkingToLeft bool) (string, string, error) {
	var leftInfixNotation string
	var leftErr error

	// 演算子が左結合性の場合、左端の単項マイナスは特別扱い
	// 例えば、-1+2 の中置表記が (-1)+2 ではなく -1+2 とならなければならない
	//
	// 左結合性でない場合、例えば (-1)^2 の中置表記は (-1)^2 のままとなる
	leftUMinus, uMinus := node.Left().(*ast.UnaryMinus)
	if uMinus && walkingToLeft && node.IsLeftAssociative() {
		leftInfixNotation, leftErr = infixNotationOfPrefixExpression(leftUMinus, walkingToLeft)
	} else {
		leftInfixNotation, leftErr = parenthesizeChildOfInfixExpression(
			node,
			node.Left(),
			node.IsLeftAssociative(),
			walkingToLeft,
		)
	}
	if leftErr != nil {
		return "", "", leftErr
	}

	rightInfixNotation, rightErr := parenthesizeChildOfInfixExpression(
		node,
		node.Right(),
		node.IsRightAssociative(),
		false,
	)
	if rightErr != nil {
		return "", "", rightErr
	}

	return leftInfixNotation, rightInfixNotation, nil
}

// infixNotationOfSumRollResult は加算ロール結果の中置表記を返す。
func infixNotationOfSumRollResult(node *ast.SumRollResult, _ bool) (string, error) {
	dieValueStrs := []string{}

	for _, d := range node.Dice {
		dieValueStrs = append(dieValueStrs, fmt.Sprintf("%d", d.Value))
	}

	return fmt.Sprintf("%d[%s]", node.Value(), strings.Join(dieValueStrs, ",")), nil
}
