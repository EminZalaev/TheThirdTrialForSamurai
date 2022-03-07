package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и
возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {

	response, err := ntp.Query("ntp1.stratum2.ru")
	if err != nil {
		log.Fatalln(err)
	}

	time := time.Now().Add(response.ClockOffset)
	hour, min, sec := time.Clock()

	fmt.Printf("Current time is : %d:%d:%d\n", hour, min, sec)
}
