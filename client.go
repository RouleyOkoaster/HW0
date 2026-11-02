package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Запрашиваем у ОС клиентский TCP-сокет и подключаемся к удаленному адресу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer conn.Close() 

	fmt.Println("Сокет успешно открыт, соединение установлено!")

	buffer := make([]byte, 1024)
	
	n, err := conn.Read(buffer) // Читаем ответ от сервера
	answer := string(buffer[:n]) // Преобразуем байты в строку
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		os.Exit(1)
	} else if answer=="OK\n" {
		fmt.Println("Ответ от сервера:", answer)
	} else {
		fmt.Println("Ответ от сервера не соответствует ожиданиям:", answer)
	}
}
