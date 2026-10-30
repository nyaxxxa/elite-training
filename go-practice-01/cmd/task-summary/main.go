package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "in_progress":
		return "В процессе"
	case "new":
		return "Новый"
	case "done":
		return "Выполнено"
	default:
		return "Статус не известен"
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	count := 0
	for i := 0; i < len(statuses); i++ {
		if statuses[i] == target {
			count++
		}
	}
	return count
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "in_progress")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)

	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Статистика: \n")
	for status, count := range counts {
		fmt.Printf("%s = %d, ", status, count)
	}
}
