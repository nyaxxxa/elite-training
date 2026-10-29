package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "in_progress":
		fmt.Println("В процессе")
	case "new":
		fmt.Println("Новый")
	case "done":
		fmt.Println("Выполнено")
	default:
		fmt.Println("Статус не известен")
	}

	return ""
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	// TODO: пройдите по срезу через for range.
	// TODO: сравните каждый элемент с target и верните счётчик.
	return 0
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	// TODO: добавьте ещё один статус new через append.

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	// TODO: создайте map[string]int для подсчёта статусов.
	// TODO: выведите new, in_progress, done в фиксированном порядке.
}
