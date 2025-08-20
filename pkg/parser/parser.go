package parser

import (
	"fmt"
	"strconv"
)

func ParseFloat(str string) (float64, error) {
	num1, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("Ошибка: введите корректное число")
	}

	return num1, nil
}
