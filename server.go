package main

import (
	"fmt"
	"net"
	"os"
)

func receiver(c net.Conn) {
	defer c.Close()
	fmt.Println("Подключение установлено с адресом:", c.RemoteAddr(), " и туда отправляется ответ 'OK'")
	c.Write([]byte("OK\n"))
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка открытия сокета:", err)
		os.Exit(1)
	}
	
	defer listener.Close()
	fmt.Println("Сокет успешно открыт. Ожидание подключений на порту 8080...")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии подключения:", err)
			break
		}

		go receiver(conn)
	}

	
}
