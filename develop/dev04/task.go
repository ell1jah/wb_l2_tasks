package main

import (
	"sort"
	"strings"
)

func MySearchAnagram(str []string) *map[string][]string {
	result := make(map[string][]string) // мапа дял результата
	helper := make(map[string]string)   // мапа для помощи поиска анаграмм
	for _, elem := range str {          // итерируемся по переданному срезу строк
		elem := strings.ToLower(elem)                                   // переводим каждую строку в нижний регистр по условию задания
		arr := []rune(elem)                                             // преобразуем строку в срез рун
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] }) // сортируем элементы среза руны по возрастаю
		if _, ok := helper[string(arr)]; !ok {                          // проверяем если это не анаграмма
			helper[string(arr)] = elem // добавляем отсортированную строку как ключ, а значение первое встретившееся слово(
			// это будет ключом для результирующей мапы)
			result[elem] = append(result[elem], elem) // добавляем впервые встретившуюся строку как ключ множества и в слайс этого множества
		} else { // если такая анаграмма уже есть в мапе помощник
			key := helper[string(arr)]              // из мапы помощника забираем значение, это будет ключом в результирующей мапе
			result[key] = append(result[key], elem) // добавляем анаграмму в срез этого множество
		}
	}
	for key, arrays := range result { //
		if len(arrays) <= 1 || key == "" { // если в срезе множества только один элемент или пустая строка, значит это не множество анаграмм
			delete(result, key) // удаляем
		} else {
			sort.Strings(arrays) // сортируем по возрастанию элементы множества
		}
	}
	return &result // возвращаем ссылку на мапу по зааднию
}
