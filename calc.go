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

func DeleteEmptySlice(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func checkNums(textArr []string) ([]string, bool, string) {
	latinNums := [9]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	arabicNums := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
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
		err := errors.New("Строка меньше 3 символов")
		return err
	}

	if len([]rune(text)) > 10 {
		err := errors.New("Слишком большая строка")
		return err
	}

	regularCheck := regexp.MustCompile(`^(?:[1-9]|I|II|III|IV|V|VI|VII|VIII|IX)[+\-*/](?:[1-9]|I|II|III|IV|V|VI|VII|VIII|IX)$`)
	if !(regularCheck.MatchString(text)) {
		err := errors.New("Строка не соотвествует формату 1 + 1, I + IV...")
		return err
	}

	textArr := []string{}
	textArr = regexp.MustCompile("[+\\-/*]").Split(text, -1)
	textArr = DeleteEmptySlice(textArr)

	textArr, operationPossible, numsType := checkNums(textArr)

	if !operationPossible {
		err := errors.New("Операция невозможна. Латинские и арабские цифры")
		return err
	}

	result := Operation(text, textArr)

	if numsType == "latin" && result < 0 {
		err := errors.New("В Римской системе нет цифр меньше 0")
		return err
	}

	fmt.Println(result)

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
