// ./mygrep [-flag (-A-B-C-c-i-v-F-n)] [pattern] [file]
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

type flags struct {
	A, B, C       int
	c, i, v, F, n bool
}

var fl flags

func init() {
	flag.IntVar(&fl.A, "A", 0, `-A - "after" печатать +N строк после совпадения`)
	flag.IntVar(&fl.B, "B", 0, `-B - "before" печатать +N строк до совпадения`)
	flag.IntVar(&fl.C, "C", 0, `-C - "context" (A+B) печатать ±N строк вокруг совпадения`)
	flag.BoolVar(&fl.c, "c", false, `-c - "count" (количество строк)`)
	flag.BoolVar(&fl.i, "i", false, `-i - "ignore-case" (игнорировать регистр)`)
	flag.BoolVar(&fl.v, "v", false, `-v - "invert" (вместо совпадения, исключать)`)
	flag.BoolVar(&fl.F, "F", false, `-F - "fixed", точное совпадение со строкой`)
	flag.BoolVar(&fl.n, "n", false, `-n - "line num", напечатать номер строки`)
	flag.Parse()
	if err := ValidateFlags(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	pattern, filename, err := PatternAndValidFile()
	if err != nil {
		log.Fatal(err)
	}
	re, err := preparePattern(pattern)
	if err != nil {
		log.Fatal(err)
	}
	text, err := scanTextFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	myGrep(text, re)
}

func PatternAndValidFile() (string, string, error) {
	args := flag.Args()
	if len(args) == 0 { // проверяем указан ли паттерн и файл, если нет возвращаем ошибку
		return "", "", fmt.Errorf("ошибка: отсутствует паттерн и файл")
	}
	if len(args) == 1 { // проверяем указан ли файл, если нет возвращаем ошибку
		return "", "", fmt.Errorf("ошибка: отсутствует паттерн или файл")
	}

	pattern := args[0]                                   // первым аргументом забираем паттерн
	filename := args[1]                                  // вторым аргументом забираем имя файла
	if _, err := os.Stat(filename); os.IsNotExist(err) { // проверяем валидность этого файла, если нет возвращаем ошибку
		return "", "", fmt.Errorf("%s - такого файла не существует", filename)
	}
	return pattern, filename, nil // возвращаем название файла с которой будем работать
}

func ValidateFlags() error {
	if fl.c { // если есть флаг c (count), то игнорируются все остальные флаги
		fl = flags{}
		fl.c = true
	}
	if fl.A < 0 || fl.B < 0 || fl.C < 0 {
		return fmt.Errorf("ошибка: флаги [A],[B],[C] не могут быть отрицательными")
	}
	if fl.C > 0 { // если есть флаг C, то приравниваю их к A и B
		fl.A = fl.C
		fl.B = fl.C
	}
	return nil
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

func preparePattern(pattern string) (*regexp.Regexp, error) { // функция для подготовки регулярного выражения
	if fl.F {
		pattern = `(\Q` + pattern + `\E)` // разделители \Q и \E предназначены для цитирования литералов
	}
	if fl.i {
		pattern = `(?i)(` + pattern + `)` //  (?i) устанавливает режим игнорирования регистра
	}
	re, err := regexp.Compile(pattern) // компилируем регулярное выражение по нашему паттерну
	if err != nil {
		return nil, err
	}
	return re, nil
}

func myGrep(text []string, re *regexp.Regexp) {
	var count int                            // переменная счетчик для подсчета строк с совпадениями
	checkprintline := make(map[int]struct{}) // счетчик считанных линий, для избежания вывода дублирования строк
	for i, line := range text {
		match := re.MatchString(line) // проверяем на содержание в строке соответствий с регулярным выражением
		if fl.v {                     // если установлен флаг инвертирования
			match = !match // вместо совпадения, будем исключать
		}
		if match { // если найдено соответствие с регуляркой
			if fl.c { // если установлен флаг подсчета, увеличиваем при каждом соответствии
				count++
			} else {
				if fl.n && fl.A == 0 && fl.B == 0 && fl.C == 0 { // если установлен флаг вывода номера строки
					fmt.Printf("%d:%s\n", i+1, line) // печатаем строку с нумерацией
				}
				if fl.B > 0 { // если установлен флаг "before"
					printABCFlag(text, 'B', i, checkprintline) // вызов функции для вывода флагов А,В,С
				}

				if fl.A > 0 { // если установлен флаг "after"
					printABCFlag(text, 'A', i, checkprintline) // вызов функции для вывода флагов А,В,С
				}

				if fl.C > 0 { // если установлен флаг "context"
					printABCFlag(text, 'C', i, checkprintline) // вызов функции для вывода флагов А,В,С
				}
			}
		}
	}
	if fl.c {
		fmt.Println(count) // выводим количество найденных строк
	}
}

func printABCFlag(text []string, flagg byte, index int, checkprintline map[int]struct{}) {
	switch flagg {
	case 'B': // если падаем в befoe
		j := index - fl.B // j переменная для итерации
		if j < 0 {
			j = 0 // если j получилась меньше 0, то приравниваем к 0 чтобы не выйти за границы слайса
		}
		for ; j <= index; j++ { // итерируемся от j к найденному индексу элемента
			if _, ok := checkprintline[j]; !ok { // проверяем по мапе если уже была такая линия то пропускаем
				checkprintline[j] = struct{}{}
				if fl.n {
					fmt.Printf("%d:%s\n", j+1, text[j]) // вывод с нумерацией строки
				} else {
					fmt.Println(text[j]) // вывод без нумерации строки
				}
			}
		}
	case 'A': // алгоритм такой же что и в B только в другую сторону (выводим элементы спереди)
		j := index + fl.A
		if j >= len(text) {
			j = len(text) - 1
		}
		for ; index <= j; index++ {
			if _, ok := checkprintline[index]; !ok {
				checkprintline[index] = struct{}{}
				if fl.n {
					fmt.Printf("%d:%s\n", index+1, text[index])
				} else {
					fmt.Println(text[index])
				}
			}
		}
	case 'C':
		printABCFlag(text, 'B', index, checkprintline) // рекурсивно вызываем с флагом B
		printABCFlag(text, 'A', index, checkprintline) // рекурсивно вызываем с флагом A
	}
}
