package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	resp, err := http.Get(os.Args[1])
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Use valid site")
	}

	recordFile("out.html", resp)
}

func recordFile(file string, m *http.Response) {
	outFile, err := os.Create(file)
	defer outFile.Close()
	if err != nil {
		log.Print("Cannot create file")
	}

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Println("Cannot read site body")
	}
	outFile.WriteString(string(body))
}
