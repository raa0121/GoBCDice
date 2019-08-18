/*
BCDiceコマンドの評価処理のパッケージ。

主な処理は、以下の3つ。

1. ダイスロールなどの可変ノードの引数を評価し、整数に変換する。

2. 可変ノードの値を決定する。

3. 抽象構文木を評価し、値のオブジェクトに変換する。

*/
package evaluator

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/object"
	"math"
)

// 評価器の構造体。
type Evaluator struct {
	diceRoller *roller.DiceRoller
	env        *Environment
}

// NewEvaluator は新しい評価器を返す。
//
// diceRoller: ダイスローラー,
// env: 評価環境
func NewEvaluator(diceRoller *roller.DiceRoller, env *Environment) *Evaluator {
	return &Evaluator{
		diceRoller: diceRoller,
		env:        env,
	}
}

// RolledDice はダイスロール結果を返す。
func (e *Evaluator) RolledDice() []dice.Die {
	return e.env.RolledDice()
}

// Eval はnodeを評価してObjectに変換し、返す。
func (e *Evaluator) Eval(node ast.Node) (object.Object, error) {
	// 型で分岐する
	switch n := node.(type) {
	case *ast.BRollList:
		return e.evalBRollList(n)
	case *ast.BRollComp:
		return e.evalBRollComp(n)
	case *ast.RRollList:
		return e.evalRRollList(n)
	case *ast.RRollComp:
		return e.evalRRollComp(n)
	case *ast.Choice:
		return e.evalChoice(n)
	case ast.Command:
		return e.Eval(n.Expression())
	case ast.PrefixExpression:
		return e.evalPrefixExpression(n)
	case ast.InfixExpression:
		return e.evalInfixExpression(n)
	case *ast.Int:
		return object.NewInteger(n.Value), nil
	case *ast.SumRollResult:
		return object.NewInteger(n.Value()), nil
	}

	return nil, fmt.Errorf("unknown type: %s", node.Type())
}

// evalBRollList はバラバラロール列を評価する。
func (e *Evaluator) evalBRollList(node *ast.BRollList) (*object.Array, error) {
	elements := []object.Object{}

	for _, b := range node.BRolls {
		o, err := e.Eval(b)
		if err != nil {
			return nil, err
		}

		intObjs := o.(*object.Array)
		elements = append(elements, intObjs.Elements...)
	}

	return object.NewArrayByMove(elements), nil
}

// evalBRollComp はバラバラロールの成功数カウントを評価する。
func (e *Evaluator) evalBRollComp(node *ast.BRollComp) (*object.BRollCompResult, error) {
	compareNode := node.Expression().(*ast.Compare)

	// 左辺を評価する
	valuesObj, evalBRollListErr := e.Eval(compareNode.Left())
	if evalBRollListErr != nil {
		return nil, evalBRollListErr
	}

	// 右辺を評価する
	evaluatedTargetObj, evalTargetErr := e.Eval(compareNode.Right())
	if evalTargetErr != nil {
		return nil, evalTargetErr
	}

	valuesArray := valuesObj.(*object.Array)
	evaluatedTargetNode :=
		ast.NewInt(evaluatedTargetObj.(*object.Integer).Value)

	// 振られた各ダイスに対して成功判定を行い、成功数を数える
	numOfSuccesses := 0
	for _, el := range valuesArray.Elements {
		valueNode := ast.NewInt(el.(*object.Integer).Value)
		valueCompareNode := ast.NewCompare(
			valueNode,
			compareNode.Operator(),
			evaluatedTargetNode,
		)

		r, compErr := e.Eval(valueCompareNode)
		if compErr != nil {
			return nil, compErr
		}

		if success := r.(*object.Boolean).Value; success {
			numOfSuccesses++
		}
	}

	return object.NewBRollCompResult(valuesArray, object.NewInteger(numOfSuccesses)), nil
}

// evalRRollList は個数振り足しロール列を評価する。
//
// TODO: シャドウラン4版のグリッチに対応する。
func (e *Evaluator) evalRRollList(node *ast.RRollList) (*object.Array, error) {
	if node.Threshold.IsNil() {
		return nil, fmt.Errorf("evalRRollList: threshold is nil")
	}

	// 閾値を評価する
	thresholdObj, evalThresholdErr := e.Eval(node.Threshold)
	if evalThresholdErr != nil {
		return nil, evalThresholdErr
	}
	thresholdInt := thresholdObj.(*object.Integer)
	threshold := thresholdInt.Value

	// ダイスロールのキュー。
	// ダイスロール後、条件を満たす出目が得られるたびに、このキューに
	// そのダイスロールを追加する。
	rollQueue := make([]*ast.RRoll, len(node.RRolls))
	copy(rollQueue, node.RRolls)

	// 最大ロール数
	// TODO: ダイスボットの設定で変更できるようにする
	maxRollCount := 1000

	// ダイスロール結果を格納する配列
	valueGroups := []object.Object{}
	for i := 0; i < maxRollCount && len(rollQueue) > 0; i++ {
		// キューの最初のダイスロールを取り出す
		rRoll := rollQueue[0]
		if len(rollQueue) < 2 {
			rollQueue = nil
		} else {
			rollQueue = rollQueue[1:len(rollQueue)]
		}

		// ダイスロールを行う
		o, err := e.Eval(rRoll)
		if err != nil {
			return nil, err
		}

		// 出目を結果の配列に格納する
		values := o.(*object.Array)
		valueGroups = append(valueGroups, values)

		// 成功数を数える
		numOfSuccesses := 0
		for _, v := range values.Elements {
			// TODO: 演算子 ">=" 以外も振り足しの条件に設定できるようにする
			if vi := v.(*object.Integer); vi.Value >= threshold {
				numOfSuccesses++
			}
		}

		// 成功した分だけ追加のダイスロールをキューに追加する
		if numOfSuccesses > 0 {
			numNode := ast.NewInt(numOfSuccesses)
			sidesNode := rRoll.Right()
			newRRoll := ast.NewRRoll(numNode, sidesNode)

			rollQueue = append(rollQueue, newRRoll)
		}
	}

	return object.NewArrayByMove(valueGroups), nil
}

// evalRRollComp は個数振り足しロールの成功数カウントを評価する。
func (e *Evaluator) evalRRollComp(node *ast.RRollComp) (*object.RRollCompResult, error) {
	compareNode := node.Expression().(*ast.Compare)

	// 左辺を評価する
	valueGroupsObj, evalRRollListErr := e.Eval(compareNode.Left())
	if evalRRollListErr != nil {
		return nil, evalRRollListErr
	}

	// 右辺を評価する
	evaluatedTargetObj, evalTargetErr := e.Eval(compareNode.Right())
	if evalTargetErr != nil {
		return nil, evalTargetErr
	}

	valueGroupsArray := valueGroupsObj.(*object.Array)
	evaluatedTargetNode :=
		ast.NewInt(evaluatedTargetObj.(*object.Integer).Value)

	// 振られた各ダイスに対して成功判定を行い、成功数を数える
	numOfSuccesses := 0
	for _, vg := range valueGroupsArray.Elements {
		valuesArray := vg.(*object.Array)
		for _, el := range valuesArray.Elements {
			valueNode := ast.NewInt(el.(*object.Integer).Value)
			valueCompareNode := ast.NewCompare(
				valueNode,
				compareNode.Operator(),
				evaluatedTargetNode,
			)

			r, compErr := e.Eval(valueCompareNode)
			if compErr != nil {
				return nil, compErr
			}

			if success := r.(*object.Boolean).Value; success {
				numOfSuccesses++
			}
		}
	}

	return object.NewRRollCompResult(
		valueGroupsArray,
		object.NewInteger(numOfSuccesses),
	), nil
}

// evalChoice はランダム選択を評価する。
func (e *Evaluator) evalChoice(node *ast.Choice) (*object.String, error) {
	rolledDice, err := e.RollDice(1, len(node.Items))
	if err != nil {
		return nil, err
	}

	index := rolledDice[0].Value - 1
	value := node.Items[index].Value

	return object.NewString(value), nil
}

// evalPrefixExpression は前置式を評価する。
func (e *Evaluator) evalPrefixExpression(
	node ast.PrefixExpression,
) (object.Object, error) {
	if node.Right() == nil {
		return nil, fmt.Errorf("operator %s: right is nil", node.Operator())
	}

	right, err := e.Eval(node.Right())
	if err != nil {
		return nil, err
	}

	if right == nil {
		return nil, fmt.Errorf("operator %s: evaluated right is nil", node.Operator())
	}

	if right.Type() == object.INTEGER_OBJ {
		return e.evalIntegerPrefixExpression(
			node.Operator(), right.(*object.Integer))
	}

	return nil, fmt.Errorf("operator not implemented: %s%s",
		node.Operator(), right.Type())
}

// evalIntegerPrefixExpression は整数ノードを子に持つ前置式を評価する。
func (e *Evaluator) evalIntegerPrefixExpression(
	operator string,
	right *object.Integer,
) (object.Object, error) {
	value := right.Value

	switch operator {
	case "-":
		return object.NewInteger(-value), nil
	}

	return nil, fmt.Errorf("operator not implemented: %s%s",
		operator, right.Type())
}

// evalInfixExpression は中置式を評価する。
func (e *Evaluator) evalInfixExpression(
	node ast.InfixExpression,
) (object.Object, error) {
	if node.Left() == nil {
		return nil, fmt.Errorf("operator %s: left is nil", node.Operator())
	}

	if node.Right() == nil {
		return nil, fmt.Errorf("operator %s: right is nil", node.Operator())
	}

	left, leftErr := e.Eval(node.Left())
	if leftErr != nil {
		return nil, leftErr
	}

	if left == nil {
		return nil, fmt.Errorf("operator %s: evaluated left is nil", node.Operator())
	}

	right, rightErr := e.Eval(node.Right())
	if rightErr != nil {
		return nil, rightErr
	}

	if right == nil {
		return nil, fmt.Errorf("operator %s: evaluated right is nil", node.Operator())
	}

	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		leftInteger := left.(*object.Integer)
		rightInteger := right.(*object.Integer)

		switch n := node.(type) {
		case ast.Divide:
			return e.evalIntegerDivide(n, leftInteger, rightInteger)
		default:
			return e.evalIntegerInfixExpression(n.Operator(), leftInteger, rightInteger)
		}
	}

	return nil, fmt.Errorf("operator not implemented: %s %s %s",
		left.Type(), node.Operator(), right.Type())
}

// evalIntegerInfixExpression は整数ノード同士を子に持つ中置式を評価する。
func (e *Evaluator) evalIntegerInfixExpression(
	operator string,
	left *object.Integer,
	right *object.Integer,
) (object.Object, error) {
	leftValue := left.Value
	rightValue := right.Value

	switch operator {
	case "+":
		return object.NewInteger(leftValue + rightValue), nil
	case "-":
		return object.NewInteger(leftValue - rightValue), nil
	case "*":
		return object.NewInteger(leftValue * rightValue), nil
	case "D":
		return e.evalSumRoll(left, right)
	case "B", "R":
		return e.evalBasicRoll(left, right)
	case "...":
		return e.evalRandomNumber(left, right)
	case "=":
		return object.NewBoolean(leftValue == rightValue), nil
	case "<>":
		return object.NewBoolean(leftValue != rightValue), nil
	case "<":
		return object.NewBoolean(leftValue < rightValue), nil
	case ">":
		return object.NewBoolean(leftValue > rightValue), nil
	case "<=":
		return object.NewBoolean(leftValue <= rightValue), nil
	case ">=":
		return object.NewBoolean(leftValue >= rightValue), nil
	}

	return nil, fmt.Errorf("operator not implemented: %s %s %s",
		left.Type(), operator, right.Type())
}

// evalIntegerDivide は除算を評価する。
func (e *Evaluator) evalIntegerDivide(
	divide ast.Divide,
	left *object.Integer,
	right *object.Integer,
) (object.Object, error) {
	leftValue := left.Value
	rightValue := right.Value

	if rightValue == 0 {
		return nil, fmt.Errorf("%d divided by zero", leftValue)
	}

	switch divide.RoundingMethod() {
	case ast.ROUNDING_METHOD_ROUND_DOWN:
		// 除算（小数点以下切り捨て）
		return object.NewInteger(leftValue / rightValue), nil
	case ast.ROUNDING_METHOD_ROUND:
		{
			// 除算（小数点以下四捨五入）
			resultFloat := math.Round(float64(leftValue) / float64(rightValue))
			return object.NewInteger(int(resultFloat)), nil
		}
	case ast.ROUNDING_METHOD_ROUND_UP:
		{
			// 除算（小数点以下切り上げ）
			resultFloat := math.Ceil(float64(leftValue) / float64(rightValue))
			return object.NewInteger(int(resultFloat)), nil
		}
	default:
		return nil, fmt.Errorf("evalIntegerDivide: unknown rounding method")
	}
}

// evalSumRoll は加算ロールを評価する。
func (e *Evaluator) evalSumRoll(
	num *object.Integer,
	sides *object.Integer,
) (object.Object, error) {
	numVal := num.Value
	sidesVal := sides.Value

	rolledDice, err := e.RollDice(numVal, sidesVal)
	if err != nil {
		return nil, err
	}

	sum := 0
	for _, d := range rolledDice {
		sum += d.Value
	}

	return object.NewInteger(sum), nil
}

// evalBasicRoll はバラバラロールを評価する。
// 返り値は、整数オブジェクトを要素として持つ配列オブジェクト、およびエラー。
func (e *Evaluator) evalBasicRoll(
	num *object.Integer,
	sides *object.Integer,
) (*object.Array, error) {
	numVal := num.Value
	sidesVal := sides.Value

	rolledDice, err := e.RollDice(numVal, sidesVal)
	if err != nil {
		return nil, err
	}

	intObjs := make([]object.Object, 0, len(rolledDice))
	for _, d := range rolledDice {
		intObjs = append(intObjs, object.NewInteger(d.Value))
	}

	return object.NewArrayByMove(intObjs), nil
}

// evalRandomNumber はランダム数値取り出しを評価する。
func (e *Evaluator) evalRandomNumber(
	min *object.Integer,
	max *object.Integer,
) (object.Object, error) {
	minValue := min.Value
	maxValue := max.Value

	if minValue >= maxValue {
		return nil,
			fmt.Errorf(
				"evalRandomNumber: min (%d) must be less than max (%d)",
				minValue,
				maxValue,
			)
	}

	randRange := maxValue - minValue + 1
	rolledDice, err := e.RollDice(1, randRange)
	if err != nil {
		return nil, err
	}

	resultValue := (minValue - 1) + rolledDice[0].Value

	return object.NewInteger(resultValue), nil
}

// RollDice は、sides個の面を持つダイスをnum個振り、その結果を返す。
// また、ダイスロールの結果を記録する。
func (e *Evaluator) RollDice(num int, sides int) ([]dice.Die, error) {
	rolledDice, err := e.diceRoller.RollDice(num, sides)
	if err != nil {
		return nil, err
	}

	e.env.AppendRolledDice(rolledDice)

	return rolledDice, nil
}
