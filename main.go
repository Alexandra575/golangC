package main

import (
	"fmt"
	"os"
	"strconv"
)

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var convRoman = [14]int{
	100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
}

func main() {
	fmt.Println("Добрый день!")

	fmt.Println("Введите выражение,используя целые арабские либо римские числа в диапазоне от 1 до 10 включительно. Между числами укажижите тип операции: +, -, * или /. После каждого символа используйте пробел.")

	var number string
	var number2 string
	var sign string
	var err = ""
	fmt.Fscanln(os.Stdin, &number, &sign, &number2, &err)
	if err != "" {
		panic("Ошибка")
		return
	}

	var numberA int
	var number2B int

	_, ok := roman[number]
	if !ok {
		numberA, _ = strconv.Atoi(number)
		number2B, _ = strconv.Atoi(number2)
		if numberA <= 0 || numberA > 10 || number2B <= 0 || number2B > 10 {
			panic("Ошибка")
			return
		} else {
			switch sign {
			case "/":
				fmt.Println(numberA / number2B)
			case "+":

				fmt.Println(numberA + number2B)
			case "-":

				fmt.Println(numberA - number2B)
			case "*":

				fmt.Println(numberA * number2B)
			default:
				fmt.Println("Ошибка, введена  недопустимая операция: +, -, *, /")

			}
		}

	} else {
		_, ok1 := roman[number2]
		if !ok1 {
			panic("Ошибка")
			return
		} else {
			roman1 := roman[number]
			roman2 := roman[number2]
			if roman1 <= 0 || roman1 > 10 || roman2 <= 0 || roman2 > 10 {
				panic("Ошибка")
				return
			} else {
				var resalt int
				switch sign {
				case "/":
					resalt = roman1 / roman2
				case "+":

					resalt = roman1 + roman2
				case "-":

					resalt = roman1 - roman2
				case "*":
					resalt = roman1 * roman2
				default:
					fmt.Println("Ошибка, введена  недопустимая операция: +, -, *, /")
				}
				var romanResalt string
				if resalt <= 0 {
					panic("Ошибка")
				} else if resalt > 0 {
					for _, numbers := range convRoman {
						for i := numbers; i <= resalt; {
							for key, value := range roman {
								if value == numbers {
									romanResalt = romanResalt + key
									resalt = resalt - numbers
								}

							}
						}

					}
				}
				fmt.Println(romanResalt)
			}
		}
	}
}
