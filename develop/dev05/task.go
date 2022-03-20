package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// ключи отвечающие за количество обработанных строк вокруг
	after := flag.Int("A", 0, "after key")
	before := flag.Int("B", 0, "before key")
	context := flag.Int("C", 0, "context key")

	//общее количество строк, которые выводят
	count := flag.Int("c", 0, "count key")

	//ключи отвечающие за обработку строк
	ignoreCase := flag.Bool("i", false, "ignore-case key")
	inverter := flag.Bool("v", false, "invertor key")
	fixed := flag.Bool("fixed", false, "fixed key")
	lineNum := flag.Bool("n", false, "live num key")
	flag.Parse()

	lines := struct {
		left  int
		rigth int
		count int
	}{left: 0, rigth: 0, count: 0}

	temp := "Клетки"

	sliceFile := handlerFile("in.txt")

	// если общее количество не указано , то возвращаем все строки
	if *count != 0 {
		lines.count = *count
	} else {
		lines.count = len(sliceFile)
	}
	// используем если указан хотя бы 1 параметр
	switch {
	case *context != 0:
		lines.left, lines.rigth = *context, *context
	case *after != 0:
		lines.rigth = *after
	case *before != 0:
		lines.left = *before

		//если не указан, то возвращаться будут все строки
	default:
		lines.rigth = len(sliceFile)
	}

	// проверяем все строки
	for i, str := range sliceFile {

		// если строка подходит - учитываются ключи обработки строки
		if handlertStr(str, temp, *ignoreCase, *inverter, *fixed) {

			// отсекаем строки, не входящие в срез
			maxOfMin := math.Max(0, float64(i-lines.left))
			minOfMax := math.Min(float64(len(sliceFile)-1), float64(i+lines.rigth))

			// захват допольнительных строк, находящихся около нужной строки
			for j := maxOfMin; j <= minOfMax; j++ {

				if lines.count >= 1 {
					// печатаем номер строки или нет - зависит от ключа
					if *lineNum {
						fmt.Println("num: ", j, " str: ", sliceFile[int(j)])
						lines.count--
					} else {
						fmt.Println(" str: ", sliceFile[int(j)])
						lines.count--
					}
				} else {
					// lines.count >= 1  количество строк для вывода
					os.Exit(1)
				}

			}
		}

	}

}

// считываем файл
func handlerFile(nameFile string) []string {
	file, _ := os.Open(nameFile)
	defer file.Close()

	byteFile, _ := io.ReadAll(file)
	stringFile := string(byteFile)
	return strings.Split(stringFile, "\n")
}

// обрабатывает строку согласно ключам
func handlertStr(str string, temp string, ignoreCase bool, invert bool, fixed bool) bool {

	// если активен ключ игнорирования регистра - приводит обе строки к одному
	if ignoreCase {
		str = strings.ToLower(str)
		temp = strings.ToLower(temp)
	}

	// если активен ключ фиксированной строки, то возвращаем зачения учитывая ключ отрицания
	if fixed {
		if str == temp {
			return true && !invert
		}
		return false || invert

	}

	// производим поиск подстроки и возвращаем зачения учитывая ключ инвентора
	if strings.Contains(str, temp) {
		return true && !invert
	}
	return false || invert

}
