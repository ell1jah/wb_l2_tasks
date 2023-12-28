package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	k       int
	n, r, u bool
}

var fl flags

func init() {
	flag.IntVar(&fl.k, "k", 0, "-k — указание колонки для сортировки ")
	flag.BoolVar(&fl.n, "n", false, "-n — сортировать по числовому значению")
	flag.BoolVar(&fl.r, "r", false, "-r — сортировать в обратном порядке")
	flag.BoolVar(&fl.u, "u", false, "-u — не выводить повторяющиеся строки")
}

func main() {
	flag.Parse()
	if fl.k < 0 {
		log.Fatal("k не может быть отрицательным")
	}
	filename, err := checkValidFile()
	if err != nil {
		log.Fatal(err)
	}
	text, err := scanTextFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	tt := mySort(text)
	for _, elem := range tt {
		fmt.Println(elem)
	}
}

func checkValidFile() (string, error) {
	args := flag.Args()
	if len(args) == 0 { // проверяем указан ли файл для сортировки , если нет возвращаем ошибку
		return "", fmt.Errorf("отсутствует файл")
	}
	filename := args[0]                                  // забираем только первое название файла
	if _, err := os.Stat(filename); os.IsNotExist(err) { // проверяем валидность этого файла , если нет возвращаем ошибку
		return "", fmt.Errorf("%s - такого файла не существует", filename)
	}
	return filename, nil // возвращаем название файла с которой будем работать
}

func scanTextFile(filename string) ([]string, error) { // функция для считывания текста из файла
	var res []string
	file, err := os.Open(filename) // открываем файл
	if err != nil {                // если есть ошибка возвращаем nil строку и ошибку
		return nil, err
	}
	defer file.Close() // отложенный вызов функция закрытия открытого файла

	scanner := bufio.NewScanner(file) // создаем сканер на базе файла file
	for scanner.Scan() {              // каждый вызов метода Scan
		// начитывает и запоминает внутри сканера очередной токен — порцию данных до разделителя(по умолчанию \n)
		text := scanner.Text()
		res = append(res, text) // добавляем в слайс
	}
	return res, scanner.Err() // возвращаем наш слайс строк, и ошибку scanner.Err если такая возникла
}

func removeDuplicates(text []string) []string { //функция дял удаления дубликатов
	var res []string
	unique := make(map[string]struct{}) // создаем мапу чтобы хранить уникальные значения
	for _, elem := range text {         // итерируемся по слайсу  строк
		if _, ok := unique[elem]; !ok { // смотрим есть ли в мапе этот элемент, если нет то добавляем в слайс и в мап
			unique[elem] = struct{}{}
			res = append(res, elem)
		}
	}
	return res // возвращаем новый слайс с удаленными дубликатами
}

func reverseFunc(text []string) {
	for i := 0; i < len(text)/2; i++ {
		text[i], text[len(text)-1-i] = text[len(text)-1-i], text[i] // в цикле переворачиваем элементы
	}
}

func sortHelpFunc(i, j int, text []string) bool { // функция помощник
	re := regexp.MustCompile(`^\d+`) // регулярное выражение для поиска чисел из начала строки
	if !fl.n {                       // если флага n нет
		str1 := strings.Split(text[i], " ")       // делим строку на слайс подстрок по пробелу
		str2 := strings.Split(text[j], " ")       // делим строку на слайс подстрок по пробелу
		if fl.k > len(str1) || fl.k > len(str2) { // проверка чтобы не уйти за границы слайса
			return false
		}
		return str1[fl.k-1] < str2[fl.k-1]
	} else { // иначе если есть флаг n
		str1 := strings.Split(text[i], " ")       // делим строку на слайс подстрок по пробелу
		str2 := strings.Split(text[j], " ")       // делим строку на слайс подстрок по пробелу
		if fl.k > len(str1) || fl.k > len(str2) { // проверка чтобы не уйти за границы слайса
			return false
		}
		num1, _ := strconv.Atoi(re.FindString(str1[fl.k-1])) // забираем число из подстроки 1
		num2, _ := strconv.Atoi(re.FindString(str2[fl.k-1])) // забираем число из подстроки 2
		return num1 < num2
	}
	return false
}

func mySort(text []string) []string {
	if fl.k == 0 && !fl.
		n { // если флаги k и n отсутствуют производим обычную сортировку с помощью стандартным методом сортировки
		sort.Slice(text, func(i, j int) bool { return text[i] < text[j] })
	}
	if fl.k > 0 {
		sort.Slice(text, func(i, j int) bool { return sortHelpFunc(i, j, text) }) //сортируем с функцией помощником
	}
	if fl.n && fl.k == 0 {
		re := regexp.MustCompile(`^\d+`) // регулярное выражение для поиска чисел из начала строки
		sort.Slice(
			text, func(i, j int) bool {
				num1, _ := strconv.Atoi(re.FindString(text[i])) // забираем число из подстроки 1
				num2, _ := strconv.Atoi(re.FindString(text[j])) // забираем число из подстроки 1
				return num1 < num2
			},
		)
	}
	if fl.u {
		text = removeDuplicates(text) // если есть флаг u, удаляем дубликаты из слайса строк
	}
	if fl.r {
		reverseFunc(text) // флаг r переворачиваем строку
	}
	return text
}
