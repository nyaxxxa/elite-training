package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	// Этот блок ввода дан готовым; вычисления добавьте самостоятельно.
	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	fmt.Printf("Получено минут: %d\n", totalMinutes)

	var hours, leftMinutes int
	hours = totalMinutes / 60
	leftMinutes = totalMinutes % 60

	if leftMinutes == 0 {
		fmt.Printf("Полных часов: %d\n", hours)

	} else {
		fmt.Printf("Полное часов: %d\n", hours)
		fmt.Printf("Оставшихся минут: %d\n", leftMinutes)
	}
}
