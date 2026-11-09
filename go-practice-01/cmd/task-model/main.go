package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	// TODO: верните "#ID Название [статус], приоритет N".
	return fmt.Sprintf("%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func main() {

	task1 := Task{ID: 4, Title: "Поступить в вуз", Status: "done", Priority: 1}
	task2 := Task{ID: 5, Title: "Закончить первый семестр", Status: "in_progress", Priority: 2}
	task3 := Task{ID: 6, Title: "Приехать домой по раньше", Status: "new", Priority: 3}
	fmt.Println(task1.Title)
	fmt.Println(task2.Title)
	fmt.Println(task3.Title)

	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	found := findTaskByID(tasks, 2)
	if found != nil {
		found.Priority = 3
		found.Complete()
	}
	count := 0

	for _, task := range tasks {
		if task.Status == "done" {
			count++
		}
	}
	fmt.Printf("Завершённые задачи = %d\n", count)
	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}
