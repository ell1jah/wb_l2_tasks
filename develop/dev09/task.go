package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"os"
)

func main() {
	url, filename, err := checkAndValidateUrl()
	if err != nil {
		log.Fatal(err)
	}
	filename += ".html"
	if checkFileExists(filename) {
		log.Fatal("такой файл уже существует")
	}
	file, err := createFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := downloadHTML(url, file); err != nil {
		log.Fatal(err)
	}
}

func checkAndValidateUrl() (string, string, error) {
	if len(os.Args) < 2 { // проверка на введенный url
		return "", "", fmt.Errorf("ошибка: отсутствует URL")
	}
	url := os.Args[1]                   // присваиваем первый аргумент к переменной
	u, err := url2.ParseRequestURI(url) // проверяем на валидность url
	if err != nil || u.Scheme == "" || u.Host == "" {
		return "", "", err
	}
	return url, u.Host, nil
}

func checkFileExists(filename string) bool { // проверяем на существования файла
	if _, err := os.Stat(filename); os.IsNotExist(err) { // если файл не существует возвращаем false
		return false
	}
	return true
}

func createFile(filename string) (*os.File, error) { //функция для создания файла
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func downloadHTML(url string, file *os.File) error {
	defer file.Close()

	response, err := http.Get(url) // делаем запрос на страничку
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(file, response.Body) // сохраняем тело запроса в файл
	if err != nil {
		return err
	}
	return nil
}
