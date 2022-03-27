package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func uniqLower(in *[]string) []string {
	res := make([]string, 0, len(*in))
	u := make(map[string]bool)

	for _, i := range *in {
		if !u[i] {
			res = append(res, strings.ToLower(i))
			u[i] = true
		}
	}

	return res
}

func anagramDict(in *[]string) map[string][]string {
	if len(*in) < 2 {
		return nil
	}

	// промежуточная мапа, ключ - отсортированное слов
	buff := make(map[string][]string)

	uniqIn := uniqLower(in)
	for _, i := range uniqIn {
		sorted := []rune(i)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})

		word := string(sorted)
		buff[word] = append(buff[word], i)

	}

	res := make(map[string][]string)
	for _, words := range buff {
		if len(words) > 1 {
			sort.Strings(words)
			res[words[0]] = words
		}
	}

	return res
}

func main() {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}

	fmt.Println(input)
	fmt.Println(anagramDict(&input))
}
