Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
выведет все числа от 0 до 9 и deadlock, мы не закрыли канал, и основная горутина завершила работу

```
