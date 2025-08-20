package app

import (
	"fmt"

	"github.com/SaveliiYam/packages/internal/usecase"
)

func StartApp() {
	fmt.Println("Простой калькулятор")
	fmt.Println("Доступные операции: +, -, *, /")
	fmt.Println("Введите 'exit' для выхода")
	fmt.Println("-----------------------------")

	usecase.Calculate()

	fmt.Println("До свидания!")
}
