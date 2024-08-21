package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX",
	70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
}

func main() {
	var input string
	fmt.Print("Введите выражение: ")
	fmt.Scanln(&input)

	isRoman := false
	for roman := range romanToInt {
		if strings.Contains(input, roman) {
			isRoman = true
			break
		}
	}

	var num1, num2 int
	var operator string

	if isRoman {
		for op := range romanToInt {
			input = strings.ReplaceAll(input, op, strconv.Itoa(romanToInt[op]))
		}
	}

	_, err := fmt.Sscanf(input, "%d %s %d", &num1, &operator, &num2)
	if err != nil {
		panic("Неправильный формат ввода!")
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("Числа должны быть в диапазоне от 1 до 10 включительно!")
	}

	result := 0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль!")
		}
		result = num1 / num2
	default:
		panic("Неизвестная операция!")
	}

	if isRoman {
		if result < 1 {
			panic("Результат римского вычисления должен быть больше нуля!")
		}
		fmt.Println("Результат:", toRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}

func toRoman(num int) string {
	if num > 100 {
		panic("Результат превышает допустимые римские числа!")
	}

	roman := ""
	for num >= 100 {
		roman += "C"
		num -= 100
	}
	for num >= 90 {
		roman += "XC"
		num -= 90
	}
	for num >= 50 {
		roman += "L"
		num -= 50
	}
	for num >= 40 {
		roman += "XL"
		num -= 40
	}
	for num >= 10 {
		roman += "X"
		num -= 10
	}
	for num > 0 {
		roman += intToRoman[num]
		num -= num
	}
	return roman
}
