package storage

import postgres "skillfactory/task_30.8.1/pkg/storage/postgres"

// Интерфейс БД.
// Этот интерфейс позволяет абстрагироваться от конкретной СУБД.
// Можно создать реализацию БД в памяти для модульных тестов.
type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	TasksByLabel(int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
	UpdateTask(postgres.Task, int) (int, error)
	DeleteTask(int) (error)
	NewLabel(postgres.Label) (int, error)
	AddLabelToTask(int, int) (error)
	NewAuthor(postgres.Author) (int, error)
}