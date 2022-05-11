package numeric

import "strconv"

//Dec2bin 十进制转2进制
func Dec2bin(num uint64) string {
	s := ""
	if num == 0 {
		return "0"
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ;num > 0 ; num /= 2 {
		lsb := num % 2
		s = strconv.FormatUint(lsb,10) + s
	}
	return s
}

//Bin2dec 二进制转十进制
func Bin2dec(s string) (num uint64) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (uint64(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}