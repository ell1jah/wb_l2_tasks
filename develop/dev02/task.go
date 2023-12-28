package main

import (
	"errors"
	"strings"
	"unicode"
)

var ErrorIncorrectString = errors.New("некорректная строка")

func UnpackingStr(str string) (string, error) {
	if str == "" { // если пустая строка то возвращаем эту же строку без ошибки
		return str, nil
	}
	rstr := []rune(str) // конвертируем строку в слайс рун , чтобы в for range не перескакивать итерации
	var result strings.Builder
	var escape bool
	for idx, val := range rstr {
		if escape {
			if idx+1 < len(rstr) && rstr[idx+1] >= '0' && rstr[idx+1] <= '9' { // если  escape true
				//  и следом идет цифра то повторяем его n раз
				result.WriteString(strings.Repeat(string(rstr[idx]), int(rstr[idx+1]-'0')))
			} else {
				result.WriteString(string(val)) // иначе записываем один раз
			}
			escape = false // возвращаем  escape  в положение false
		} else if val == '\\' { // если видим  escape то переклчаем булево значение в true
			escape = true
		} else if val >= '0' && val <= '9' {
			if idx == 0 { // если первый элемент в строке цифра = некорректная строка
				return "", ErrorIncorrectString
			}
			if idx+1 < len(rstr) && rstr[idx+1] >= '0' && rstr[idx+1] <= '9' { // если две цифры идут подряд = некорректная строка
				return "", ErrorIncorrectString
			}
			if unicode.IsLetter(rstr[idx-1]) { // проверяем если элемент до цифры символ то повторяем его n раз (
				// это нужно чтобы при escpe не повторялся повтор)
				result.WriteString(strings.Repeat(string(rstr[idx-1]), int(val-'0')))
			}
		} else if idx+1 < len(rstr) && unicode.IsLetter(val) && !unicode.IsDigit(rstr[idx+1]) { // если просто символ и след эелемент не цифра, записываем в единственном экземпляре
			result.WriteString(string(val))
		}
	}
	if unicode.IsLetter(rstr[len(rstr)-1]) { // поседний случай , проверка последнего элемента на символ
		result.WriteString(string(rstr[len(rstr)-1]))
	}
	return result.String(), nil
}
