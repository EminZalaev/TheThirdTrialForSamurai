package main

import (
	"bufio"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	k := flag.Int("k", 0, "enter column")
	n := flag.Bool("n", false, "sort by numeric value")
	r := flag.Bool("r", false, "sort in reverse order")
	u := flag.Bool("u", false, "do not output duplicate lines")
	flag.Parse()

	// читаем файл и данные преобразуем в сроковый слайс
	dataFile := readingFile("in.txt")

	switch {
	case *u:
		dataFile = uniqueString(dataFile)
	case *k != 0: // кейс обрабатывает случай, если колонка была указана. так же внутри кейса проверяются и другие ключи
		dataFile = sortByColumn(dataFile, *k)

	case *r == true: // кейс обрабатывает случай если нужна только сортировка в обратном порядке
		dataFile = reversSort(dataFile)

	case *n == true: // сортировка численно
		dataFile = sortNumeric(dataFile)
	}

	recordFile("out.txt", dataFile)

}

// считывает построчно файл и возвращает слайс
func readingFile(file string) []string {
	var result []string

	inFile, _ := os.Open(file)
	defer inFile.Close()

	fileScanner := bufio.NewScanner(inFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		result = append(result, str)
	}
	return result
}

// функция раписывает данные в файл
func recordFile(file string, array []string) {
	outFile, _ := os.Create(file)
	defer outFile.Close()

	for i := 0; i < len(array)-1; i++ {
		outFile.WriteString(array[i] + "\n")
	}
	outFile.WriteString(array[len(array)-1])
}

func uniqueString(noneUniqueString []string) []string {
	for i, str := range noneUniqueString {
		for j := i + 1; j < len(noneUniqueString); j++ {
			if str == noneUniqueString[j] {
				noneUniqueString = append(noneUniqueString[:i], noneUniqueString[j:]...)
			}
		}
	}
	return noneUniqueString
}

func sortNumeric(unsorted []string) []string {

	for i := range unsorted {
		var count1 int
		l := strings.Split(unsorted[i], " ")

		for j := range l {
			_, err := strconv.Atoi(l[j])
			if err == nil {
				count1++
			}
		}

		if count1 == len(l) {
			var result []int

			for j := range l {
				k, _ := strconv.Atoi(l[j])
				result = append(result, k)
			}

			sort.Ints(result)

			for m := range result {
				l = append(l[:m], strconv.Itoa(result[m]))
			}
		}
		unsorted[i] = strings.Join(l, " ")
	}
	return unsorted
}

func reversSort(unsorted []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(unsorted)))
	return unsorted
}

func sortByColumn(data []string, colNum int) []string {
	sort.Slice(data, func(i, j int) bool {
		return data[i][colNum] < data[j][colNum]
	})

	return data
}
