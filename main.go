package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

var arabicToRoman = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToInt(roman string) int {
	n := len(roman)
	result := 0
	for i := 0; i < n; i++ {
		if i+1 < n && romanToArabic[roman[i:i+2]] != 0 {
			result += romanToArabic[roman[i:i+2]]
			i++
		} else {
			result += romanToArabic[string(roman[i])]
		}
	}
	return result
}

func intToRoman(num int) string {
	if num < 1 {
		panic("Результат вычисления < 1")
	}
	result := ""
	for _, entry := range arabicToRoman {
		for num >= entry.Value {
			result += entry.Symbol
			num -= entry.Value
		}
	}
	return result
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль")
		}
		return a / b
	default:
		panic("Неверный оператор")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите пример для вычисления ('2+2' или 'IV*VI'), или введите 'exit' для выхода:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Выход из калькулятора.")
			break
		}

		fmt.Println("Input:\n", input)

		re := regexp.MustCompile(`^\s*([IVXLCDM0-9]+)\s*([+\-*/])\s*([IVXLCDM0-9]+)\s*$`)
		matches := re.FindStringSubmatch(input)

		if len(matches) != 4 {
			panic("Введите пример в формате '2 + 6' или 'IV * X'\nИспользуйте только +, -, *, / операторы")
		}

		aStr, operator, bStr := matches[1], matches[2], matches[3]

		if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
			panic("Неверный оператор. Используйте только +, -, *, /")
		}

		var a, b int
		var isRomanInput bool

		isRomanA := regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(aStr)
		isRomanB := regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(bStr)

		if isRomanA && isRomanB {
			a = romanToInt(aStr)
			b = romanToInt(bStr)
			isRomanInput = true
		} else if !isRomanA && !isRomanB {
			var err error
			a, err = strconv.Atoi(aStr)
			if err != nil {
				fmt.Println("Неверный ввод первого числа:", aStr)
				continue
			}
			b, err = strconv.Atoi(bStr)
			if err != nil {
				fmt.Println("Неверный ввод второго числа:", bStr)
				continue
			}
			isRomanInput = false
		} else {
			panic("Нельзя смешивать римские и арабские числа")
		}

		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("Упс! Я умею делать вычисления только с числами от 1 до 10")
		}

		result := calculate(a, b, operator)

		if isRomanInput {
			fmt.Println("Output:\n", intToRoman(result))
		} else {
			fmt.Println("Output:\n", result)
		}
	}
}
