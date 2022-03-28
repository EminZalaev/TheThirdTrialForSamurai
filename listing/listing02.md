Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
т.к defer в первой функции выполнился при возврате переменной, соответственно x++ выполнился
во второй функции мы возвращаем конкретно значение, после выполняется defer
defer выполняется в конце функции. Последний defer выполнится первым

```
