package bigint

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Bigint ...
type Bigint struct {
	Value string
}

var (
	ErrorInput     = errors.New("bad input, please input only number")
	ErrorNotNumber = errors.New("input not a number")
)

func removeZeros(num string) string {
	start := 0
	check := strings.HasPrefix(num, "+")
	if check {
		num = num[1:]
	}
	for strings.HasPrefix(num, "0") && len(num) > 1 {
		num = num[1:]
	}
	if strings.HasPrefix(num, "-") {
		start = 1
		num = num[1:]
	}
	if start == 1 && num != "0" {
		num = "-" + num
	}

	return num
}

func NewInt(num string) (Bigint, error) {

	var f bool
	var err error
	if f, err = validation(num); err != nil {
		return Bigint{}, err
	}
	if !f {
		return Bigint{}, ErrorInput
	}
	num = removeZeros(num)

	return Bigint{Value: num}, nil

}

func validation(num string) (bool, error) {

	if match, err := regexp.MatchString(`^[+-]?[0-9]*$`, num); err != nil {
		return false, err
	} else {
		return match, nil
	}

}

func (z *Bigint) Set(num string) error {
	var f bool
	var err error

	if f, err = validation(num); err != nil {
		return err
	}

	if !f {
		return ErrorInput
	}
	z.Value = removeZeros(num)
	return nil
}
func compareStringsByValue(x, y string) int {

	result := 0
	if len(x) > len(y) {
		result = 1
	} else if len(x) < len(y) {
		result = -1
	} else {
		for i := 0; i < len(x); i++ {
			if int(x[i]) > int(y[i]) {
				result = 1
				break

			} else if int(x[i]) < int(y[i]) {
				result = -1
				break
			}
		}
	}
	return result
}

// Add ...
func Add(a, b Bigint) Bigint {
	x, y := a.Value, b.Value
	var num1 = strings.Split(x, "")
	var num2 = strings.Split(y, "")
	var checkPlusMinus_x = num1[0]
	var checkPlusMinus_y = num2[0]
	var answer string = ""

	val_x := trimFirstOperator(x) //

	val_y := trimFirstOperator(y)
	x = val_x
	y = val_y
	if checkPlusMinus_y == "-" && checkPlusMinus_x == "-" {
		return Bigint{Value: "-" + add(x, y)}
		// +a + (+b)
	}
	if checkPlusMinus_x == "-" && checkPlusMinus_y == "+" && compareStringsByValue(x, y) == 1 {
		answer = removeZeros("-" + sub(x, y))
		// +a + (+b)
	}
	if checkPlusMinus_x == "-" && checkPlusMinus_y == "+" && compareStringsByValue(x, y) == -1 {
		answer = removeZeros(sub(y, x))
		// +a + (-b)
	}
	if checkPlusMinus_x == "+" && checkPlusMinus_y == "-" && compareStringsByValue(x, y) == 1 {
		answer = removeZeros(sub(x, y))
		// +a + (+b)
	}
	if checkPlusMinus_x == "+" && checkPlusMinus_y == "-" && compareStringsByValue(x, y) == -1 {
		answer = "-" + removeZeros(sub(y, x))
		// +a + (+b)
	} else if checkPlusMinus_y == "+" && checkPlusMinus_x == "+" {
		answer = add(x, y)
		// -a + (-b)
	}
	return Bigint{Value: answer}
}
func add(x, y string) string {

	var len1 = len(x)
	var len2 = len(y)

	var diffLen1 = len1 - len2
	var diffLen2 = len2 - len1
	if len1 > len2 {
		for i := 0; i < diffLen1; i++ {
			y = "0" + y
		}
	}

	if len2 > len1 {
		for i := 0; i < diffLen2; i++ {
			x = "0" + x
		}
	}
	var num1 = strings.Split(x, "")
	var num2 = strings.Split(y, "")
	var answer string
	var result int
	var carry int
	for i := len(num1) - 1; i >= 0; i-- {
		first, _ := strconv.Atoi(num1[i])
		second, _ := strconv.Atoi(num2[i])
		result = first + second + carry
		remainder := result % 10
		if result >= 10 {
			answer = strconv.Itoa(remainder) + answer
			carry = result / 10

		} else {
			answer = strconv.Itoa(result) + answer
			carry = 0
		}
	}
	if carry != 0 {
		answer = strconv.Itoa(carry) + answer
	}

	return answer

}

// Sub ...
func Sub(a, b Bigint) Bigint {
	x, y := a.Value, b.Value

	var num1 = strings.Split(x, "")
	var num2 = strings.Split(y, "")
	var checkPlusMinus_x = num1[0]
	var checkPlusMinus_y = num2[0]

	val_x := trimFirstOperator(x) //

	val_y := trimFirstOperator(y)
	x = val_x
	y = val_y

	var answer string
	// a < b +a +(-b) == b-a
	if checkPlusMinus_x == "+" && checkPlusMinus_y == "-" && compareStringsByValue(x, y) == -1 {
		answer = add(y, x)
		// a > b +a + (-b)
	} else if checkPlusMinus_x == "+" && checkPlusMinus_y == "-" && compareStringsByValue(x, y) == 1 {

		answer = add(x, y)
		// -a + b a>b
	} else if checkPlusMinus_x == "-" && checkPlusMinus_y == "+" {
		answer = "-" + add(x, y)
		// -a +b a<b
	} else if checkPlusMinus_x == "+" && checkPlusMinus_y == "+" && compareStringsByValue(x, y) == -1 {
		answer = "-" + sub(y, x)
	} else if checkPlusMinus_x == "+" && checkPlusMinus_y == "+" && compareStringsByValue(x, y) == 1 {
		answer = sub(x, y)
	} else if checkPlusMinus_x == "-" && checkPlusMinus_y == "-" && compareStringsByValue(x, y) == 1 {
		answer = "-" + sub(x, y)
	}

	return Bigint{Value: answer}

}

func sub(x, y string) string {

	len1, len2 := len(x), len(y)

	difflen1 := len1 - len2
	difflen2 := len2 - len1
	if len1 > len2 {
		for i := 0; i < difflen1; i++ {
			y = "0" + y
		}
	}

	if len2 > len1 {
		for i := 0; i < difflen2; i++ {
			x = "0" + x
		}
	}
	var num1 = strings.Split(x, "")
	var num2 = strings.Split(y, "")
	var answer string
	var result int
	for i := len(x) - 1; i >= 0; i-- {
		first, _ := strconv.Atoi(num1[i])
		second, _ := strconv.Atoi(num2[i])
		if first >= second {
			result = first - second
		} else if second > first {
			beforeLast1, _ := strconv.Atoi(num1[i-1])
			num1[i-1] = strconv.Itoa(beforeLast1 - 1)
			result = 10 + first - second
		}
		answer = strconv.Itoa(result) + answer
	}
	return answer
}

// Multiply ...
func Multiply(x, y Bigint) Bigint {
	first, _ := strconv.ParseInt(x.Value, 10, 32)
	second, _ := strconv.ParseInt(y.Value, 10, 32)
	multiplication := first * second

	return Bigint{strconv.FormatInt(int64(multiplication), 10)}
}

// Mod ...
func Mod(x, y Bigint) Bigint {
	first, _ := strconv.ParseInt(x.Value, 10, 64)
	second, _ := strconv.ParseInt(y.Value, 10, 64)
	modulus := first % second
	return Bigint{strconv.FormatInt(int64(modulus), 10)}

}

// Abs ...
func (z Bigint) Abs() Bigint {
	if z.Value[0] == '-' {
		return Bigint{
			Value: z.Value[1:],
		}
	}
	if z.Value[0] == '+' {
		return Bigint{
			Value: z.Value[1:],
		}
	}
	return Bigint{
		Value: z.Value,
	}

}

func trimFirstOperator(x string) string {
	for i := range x {
		if i > 0 {
			return x[i:]
		}
	}
	return ""
}
