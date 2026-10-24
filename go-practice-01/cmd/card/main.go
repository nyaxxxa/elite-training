package main

import "fmt"

func main() {
	fmt.Println("Задание 2: здесь будет карточка участника курса.")

	const language = "Go"
	var name string = "Эрдни"
	completed := 1
	var ready bool = true

	fmt.Printf("Изучаю курс по %s\n", language)
	fmt.Printf("Участник курса %s\n", name)
	fmt.Printf("Выполнено: %d\n", completed)
	fmt.Printf("Готовность к кодингу: %t\n", ready)

}
