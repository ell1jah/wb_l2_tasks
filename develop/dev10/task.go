package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type config struct {
	timeout    time.Duration
	host, port string
}

var cfg config

func init() {
	flag.DurationVar(&cfg.timeout, "timeout", 10*time.Second, "срок ответа на запрос (задается в секундах пример: 10s)")
}

func main() {
	flag.Parse()
	if err := validateFlags(); err != nil {
		log.Fatal(err)
	}
	quit, cancel := make(chan os.Signal, 1), make(chan struct{})
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // сигналы ос

	go myTelnet(cancel)
	<-quit        // блокируем мейн ждем сигнала ос
	close(cancel) // закрываем канал cancel чтобы завершить функцию telnet
}

func validateFlags() error {
	args := flag.Args()
	if len(args) == 0 {
		return fmt.Errorf("ошибка: не указан host и port")
	}
	if len(args) == 1 {
		return fmt.Errorf("ошибка: не указан port")
	}
	cfg.host = args[0]
	cfg.port = ":" + args[1]
	return nil
}

func myTelnet(cancel chan struct{}) {
	conn, err := net.DialTimeout("tcp", cfg.host+cfg.port, cfg.timeout) //подключение к серверу по tcp с таймаутом
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() { // запуск горутины для чтения данных из сокета и вывода их в stdout
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			conn.Close()
			log.Fatal(err)
		}
	}()

	go func() { // чтение данных из stdin и отправка их в сокет
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data := scanner.Text()
			_, err := fmt.Fprintln(conn, data)
			if err != nil {
				log.Fatal(err)
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}()
	<-cancel // блокируемся и ждем сигнала
}
