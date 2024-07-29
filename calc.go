package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var latinNums = [100]string{
	"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
	"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}

var arabicNums = CreateNumbersSlice()

func CreateNumbersSlice() []string {
	numbers := make([]string, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = strconv.Itoa(i + 1)
	}
	return numbers
}

func DeleteEmptySlice(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func CheckNums(textArr []string) ([]string, bool, string) {
	operationPossible := false
	firstNumLatin := false
	secondNumLatin := false
	numsType := "arabic"

	for i, value := range latinNums {
		if textArr[0] == value {
			firstNumLatin = true
			textArr[0] = arabicNums[i]
			numsType = "latin"
		}

		if textArr[1] == value {
			secondNumLatin = true
			textArr[1] = arabicNums[i]
			numsType = "latin"
		}
	}

	if firstNumLatin == secondNumLatin {
		operationPossible = true
	}

	return textArr, operationPossible, numsType
}

func Operation(text string, textArr []string) (result int) {
	result = 0
	firstNum, _ := strconv.Atoi(textArr[0])
	secondNum, _ := strconv.Atoi(textArr[1])

	if strings.Contains(text, "+") {
		result = firstNum + secondNum
	}
	if strings.Contains(text, "-") {
		result = firstNum - secondNum
	}
	if strings.Contains(text, "/") {
		result = firstNum / secondNum
	}
	if strings.Contains(text, "*") {
		result = firstNum * secondNum
	}

	return result
}

func CheckText(text string) error {
	if text == "" {
		err := errors.New("Строка пустая")
		return err
	}

	if len([]rune(text)) < 3 {
		err := errors.New("Выдача паники, так как строка не является математической операцией.")
		return err
	}

	if len([]rune(text)) > 10 {
		err := errors.New("Слишком большая строка")
		return err
	}

	regularCheck := regexp.MustCompile(`^(?:[1-9]|10|I|II|III|IV|V|VI|VII|VIII|IX|X)[+\-*/](?:[1-9]|10|I|II|III|IV|V|VI|VII|VIII|IX|X)$`)
	if !(regularCheck.MatchString(text)) {
		err := errors.New("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return err
	}

	textArr := []string{}
	textArr = regexp.MustCompile("[+\\-/*]").Split(text, -1)
	textArr = DeleteEmptySlice(textArr)

	textArr, operationPossible, numsType := CheckNums(textArr)

	if !operationPossible {
		err := errors.New("Выдача паники, так как используются одновременно разные системы счисления.")
		return err
	}

	result := Operation(text, textArr)

	if numsType == "latin" && result < 0 {
		err := errors.New("Выдача паники, так как в римской системе нет отрицательных чисел.\n")
		return err
	}

	if numsType == "latin" && result < 1 {
		err := errors.New("Выдача паники, результат работы с римскими цифрами меньше еденицы.\n")
		return err
	}

	if numsType == "latin" {
		fmt.Println(latinNums[result-1])
	} else {
		fmt.Println(result)
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите значение")

	text, _ := reader.ReadString('\n')

	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "")

	err := CheckText(text)
	if err != nil {
		fmt.Println(err)
	}
}
