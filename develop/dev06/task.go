package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	var fields = flag.Int("f", -1, "выбрать поля (колонки)")
	var delimiter = flag.String("d", "\t", "использовать другой разделитель")
	var separated = flag.Bool("s", false, "только строки с разделителем")

	log.SetFlags(0)
	flag.Parse()

	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	delimiterByte := []byte(*delimiter)
	lines := bytes.Split(content, []byte("\n"))

	if *separated {
		separatedBytes(lines, delimiterByte)
	}

	delimiterBytes(lines, delimiterByte, fields)

}

func delimiterBytes(lines [][]byte, delimiterByte []byte, fields *int) {
	res := make([][][]byte, 0)
	for i := range lines {
		temp := bytes.Split(lines[i], delimiterByte)
		if field := *fields; field > -1 {
			if len(temp) > field {
				temp = [][]byte{temp[field]}
			} else {
				temp = [][]byte{}
			}
		}
		res = append(res, temp)
	}

	fmt.Printf("%q", res)
}

func separatedBytes(lines [][]byte, delimiterByte []byte) [][]byte {
	temp := make([][]byte, 0)
	for i := range lines {
		if bytes.Contains(lines[i], delimiterByte) {
			temp = append(temp, lines[i])
		}
	}
	lines = temp
	return lines
}
