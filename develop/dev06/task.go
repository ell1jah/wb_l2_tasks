package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	f, d   string
	s      bool
	fields []int
}

var fl flags

func init() {
	flag.StringVar(&fl.f, "f", "", `-f - "fields" - выбрать поля (колонки)`)
	flag.StringVar(&fl.d, "d", "\t", `-d - "delimiter" - использовать другой разделитель`)
	flag.BoolVar(&fl.s, "s", false, `-s - "separated" - только строки с разделителем`)
}

var (
	errorNegativeNum = errors.New("ошибка: -f (fields) не может быть нулем или отрицательным")
)

func main() {
	flag.Parse()
	if err := validFlagF(); err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	result := myCut(scanner)
	fmt.Println(result)
}

func validFlagF() error {
	if len(fl.f) == 0 {
		return fmt.Errorf("ошибка: флаг -f (fields) должен быть указан") // при отсутствии флага f возвращаем ошибку
	}
	fields := strings.Split(fl.f, ",") //разбиваем строку в срез , разделитель запятая "1,2,3" = [1 2 3]
	unique := make(map[int]struct{})   // создаем мапу чтобы хранить уникальные значения
	for _, elem := range fields {      // итерируемся по срезу полученных строк
		colNum, err := strconv.Atoi(elem) // преобразуем в инт
		if err != nil {
			return fmt.Errorf("ошибка: флаг -f (fields) должен содержать только числа и цифры > 0 -[%s]", err.Error())
		}
		if colNum <= 0 { // проверка на отрицательное значение
			return errorNegativeNum
		}
		if _, ok := unique[colNum]; !ok { // смотрим есть ли в мапе этот элемент, если нет то добавляем в слайс и в мап
			unique[colNum] = struct{}{}
			fl.fields = append(fl.fields, colNum)
		}
	}
	sort.Ints(fl.fields) // сортируем полученные значения колонок
	return nil
}

func myCut(scanner *bufio.Scanner) string {
	var resultstring strings.Builder //  результирующая строка
	//scanner := bufio.NewScanner(os.Stdin) // созадем сканер который будет считывать с stdin, по умолчанию разделитель \n
	for scanner.Scan() { // сканируем c помощью bufio.Scanner строку
		data := scanner.Text()                     // забираем строку
		if fl.s && !strings.Contains(data, fl.d) { // если есть флаг s и нету разделителя d, то пропускаем эту строку и переходим к след итерации
			continue
		}
		if !strings.Contains(data, fl.d) { // если нет разделителя d, то печатаем строку целиком и переходим к след итерации
			resultstring.WriteString(data)
			resultstring.WriteString("\n")
			continue
		}
		res := strings.Split(data, fl.d)   // делим строку по разделителю
		for _, column := range fl.fields { // итерируемся  по колонкам которые указали
			if column > len(res) { // если колонка превышает размер среза ,
				// останавливаем цикл так как след цифры отсортированы и они тоже выходят за пределы
				break
			}
			resultstring.WriteString(res[column-1]) // добавляем колонку к результату
			resultstring.WriteString(fl.d)          // добавляем разделитель
		}
		resultstring.WriteString("\n") // после каждой итерации добавляем символ новой строки
	}
	return strings.TrimSuffix(resultstring.String(), "\n")
}
