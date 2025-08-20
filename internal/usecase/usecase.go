package usecase

import (
	"errors"
	"fmt"

	"github.com/SaveliiYam/packages/pkg/calculate"
	"github.com/SaveliiYam/packages/pkg/parser"
	"github.com/SaveliiYam/packages/pkg/terminal"
)

var errExit = errors.New("exit")
var errUndefinedOperator = errors.New("undefined operator")

func getFloatOrExit(a int) (float64, error) {
	switch a {
	case 1:
		fmt.Print("Введите первое число: ")
	case 2:
		fmt.Print("Введите второе число: ")
	}

	numStr := terminal.ReadString()

	if numStr == "exit" {
		return 0, errExit
	}

	num, err := parser.ParseFloat(numStr)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func getOperator() (string, error) {
	fmt.Print("Введите оператор (+, -, *, /): ")

	operator := terminal.ReadString()

	if operator == "exit" {
		return "", errExit
	}

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Ошибка: неверный оператор")
		return "", errUndefinedOperator
	}

	return operator, nil
}

func Calculate() {
	for {
		// Получаем первое число
		num1, err := getFloatOrExit(1)
		if err != nil {
			if errors.Is(err, errExit) {
				return
			}

			fmt.Println(err.Error())
			continue
		}

		// Получаем оператор
		operator, err := getOperator()
		if err != nil {
			if errors.Is(err, errExit) {
				return
			}

			continue
		}

		// Получаем второе число
		num2, err := getFloatOrExit(2)
		if err != nil {
			if errors.Is(err, errExit) {
				return
			}

			fmt.Println(err.Error())
			continue
		}

		// Выполняем вычисление
		result, err := calculate.Calculate(num1, num2, operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Printf("Результат: %.2f\n\n", result)
	}
}
