package main

import "fmt"

type Computer interface {
	GetType() string
	PrintDetails()
}

// NewComputer централизованный конструктор создания объектов, которому передаем тип и конструктор возвращает соответсвующий объект
func NewComputer(typeName string) Computer {
	switch typeName {
	case "server":
		return NewServer()
	case "pc":
		return NewPC()
	default:
		fmt.Println("неизвестный тип")
		return nil
	}
}

type Server struct {
	Type   string
	Memory int
}

func NewServer() Computer {
	return &Server{Type: "Server", Memory: 128}
}

func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetails() {
	fmt.Printf("Type - [%s], Memory - [%d]\n", s.Type, s.Memory)
}

type PC struct {
	Type   string
	Memory int
}

func NewPC() Computer {
	return &Server{Type: "pc", Memory: 16}
}

func (p PC) GetType() string {
	return p.Type
}

func (p PC) PrintDetails() {
	fmt.Printf("Type - [%s], Memory - [%d]\n", p.Type, p.Memory)
}

func main() {
	computer1 := NewComputer("server")
	computer2 := NewComputer("pc")
	computer1.PrintDetails()
	computer2.PrintDetails()
}
