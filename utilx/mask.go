package utilx

// MaskPhone 手机号脱敏
func MaskPhone(phone string) string {
	if len(phone) < 8 {
		return "****"
	}
	return phone[:3] + "****" + phone[7:]
}

// MaskPhoneEndFront 手机号脱敏-中件替换8个 *
func MaskPhoneEndFront(phone string, appendLen int, appendStr string) string {
	if len(phone) < 2 {
		return "****"
	}
	phones := []rune(phone)
	str := ""
	for i := 0; i < appendLen; i++ {
		str = str + appendStr
	}
	return string(phones[0]) + str + string(phones[len(phones)-1])
}

// MaskNameEnd 名称脱敏 保留最后一个字
func MaskNameEnd(name string, appendLen int, appendStr string) string {
	runes := []rune(name)
	if len(runes) <= 1 {
		return "**"
	}
	str := ""
	for i := 0; i < appendLen; i++ {
		str = str + appendStr
	}
	return str + string(runes[len(runes)-1])
}

// MaskNameFront 名称脱敏 保留最后一个字
func MaskNameFront(name string) string {
	runes := []rune(name)
	if len(runes) <= 1 {
		return "**"
	}
	str := ""
	for i := 0; i < len(runes)-1; i++ {
		str = str + "*"
	}
	return string(runes[0]) + str
}
