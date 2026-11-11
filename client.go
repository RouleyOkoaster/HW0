package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Сокет успешно открыт, соединение установлено!")

	buffer, err := io.ReadAll(conn)

	answer := string(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		os.Exit(1)
	} else if answer == "OK\n" {
		fmt.Println("Ответ от сервера:", answer)
	} else {
		fmt.Fprintln(os.Stderr, "Ответ от сервера не соответствует ожиданиям: ", answer)
		os.Exit(1)
	}
}
