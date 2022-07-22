package main

import (
	"fmt"
	"log"
	"os"
	"skillfactory/task_30.8.1/pkg/storage"
	postgres "skillfactory/task_30.8.1/pkg/storage/postgres"
)

// Интерфейс БД.
var db storage.Interface

func main() {
	var err error
	os.Setenv("dbpass", "123456")

	pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}
	connstr :=
		"postgres://postgres:" + pwd + "@localhost:5432/task_30_8_1"
	// присвоение переменной типа интерфейс конкретной реализации БД
	db, err = postgres.New(connstr)
	if err != nil {
		log.Fatal(err)
	}

	//добавление нового автора
	author_id, err := db.NewAuthor(postgres.Author{Name: "Сашка"})
	if err != nil {
		log.Fatal(err)
	}

	//добавление новой метки
	label_id, err := db.NewLabel(postgres.Label{Name: "Первая задача"})
	if err != nil {
		log.Fatal(err)
	}

	//добавление новой метки к задачи 1
	err = db.AddLabelToTask(label_id, 1)
	if err != nil {
		log.Fatal(err)
	}

	//добавление задачи
	_, err = db.NewTask(postgres.Task{AuthorID: author_id, Title: "задача", Content: "задача типа того"})
	if err != nil {
		log.Fatal(err)
	}

	//вывод всех задач
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	//обновление задачи c id 1
	_,err = db.UpdateTask(postgres.Task{AuthorID: 0, Title: "обновленная задача", Content: "обновленная задача типа того", Opened: 3, Closed: 3}, 1)
	if err != nil {
		log.Fatal(err)
	}

	//удаление задачи с ИД 2
	err = db.DeleteTask(2)
	if err != nil {
		log.Fatal(err)
	}

	tasks_by_label, err := db.TasksByLabel(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)
	fmt.Println()
	fmt.Println(tasks_by_label)
}