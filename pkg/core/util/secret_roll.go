package util

// CheckIfInputMayBeASecretRoll は input がシークレットロールの可能性があるか確認する。
//
// 返り値:
// シークレットロールのマークを除去した入力文字列,
// シークレットロールの可能性があるか（true/false）
func CheckIfInputMayBeASecretRoll(input string) (string, bool) {
	if len(input) < 1 {
		return "", false
	}

	inputRunes := []rune(input)

	switch inputRunes[0] {
	case 'S', 's':
		return string(inputRunes[1:len(inputRunes)]), true
	default:
		return input, false
	}
}
