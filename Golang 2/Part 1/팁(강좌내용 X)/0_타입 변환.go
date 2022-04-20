package main

func main() {

}

// strconv 패키지의 함수

// func Atoi(s string) (i int, err error): 숫자로 이루어진 문자열을 숫자로 변환
// func Itoa(i int) string: 숫자를 문자열로 변환
// func FormatBool(b bool) string: 불 값을 문자열로 변환
// func FormatFloat(f float64, fmt byte, prec, bitSize int) string: 실수를 문자열로 변환
// func FormatInt(i int64, base int) string: 부호 있는 정수를 문자열로 변환
// func FormatUint(i uint64, base int) string: 부호 없는 정수를 문자열로 변환
// func AppendBool(dst []byte, b bool) []byte: 불 값을 문자열로 변환하여 슬라이스 뒤에 추가
// func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte: 실수를 문자열로 변환하여 슬라이스 뒤에 추가
// func AppendInt(dst []byte, i int64, base int) []byte: 부호 있는 정수를 문자열로 변환하여 슬라이스 뒤에 추가
// func AppendUint(dst []byte, i uint64, base int) []byte: 부호 없는 정수를 문자열로 변환하여 슬라이스 뒤에 추가
// func ParseBool(str string) (value bool, err error): 불 형태의 문자열을 불로 변환
// func ParseFloat(s string, bitSize int) (f float64, err error): 실수 형태의 문자열을 실수로 변환
// func ParseInt(s string, base int, bitSize int) (i int64, err error): 부호 있는 정수 형태의 문자열을 부호 있는 정수로 변환
// func ParseUint(s string, base int, bitSize int) (n uint64, err error): 부호 없는 정수 형태의 문자열을 부호 없는 정수로 변환
