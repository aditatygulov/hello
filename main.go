package main

import (
	"fmt"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

// LuhnCheck возвращает true, если cardNumber проходит проверку по
// алгоритму Луна. Пустая строка и любые нецифровые символы → false.
func LuhnCheck(cardNumber string) bool {
	if cardNumber == "" {
		return false
	}

	// Проверяем, что все символы — цифры.
	for i := 0; i < len(cardNumber); i++ {
		if cardNumber[i] < '0' || cardNumber[i] > '9' {
			return false
		}
	}

	// Идём справа налево.
	// Последнюю цифру не удваиваем.
	sum := 0
	double := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if double {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}

func DetectBank(cardNumber string, banks []Bank) *Bank {
	if cardNumber == "" {
		return nil
	}

	for i := range banks {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}
	}

	return nil
}

func main() {
	banks := []Bank{
		{Name: "Lunar Bank", Prefix: "4000"},
		{Name: "Mars", Prefix: "5000"},
		{Name: "Saturn", Prefix: "6000"},
	}

	cardNumber := "4000123456789017"

	fmt.Println("Валиден по Луне:", LuhnCheck(cardNumber))

	bank := DetectBank(cardNumber, banks)

	if bank != nil {
		fmt.Println("Банк:", bank.Name)
	} else {
		fmt.Println("Банк не найден")
	}
}
