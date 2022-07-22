package storage

import (
	"log"
	"os"
	"testing"
)

var s *Storage

func TestMain(m *testing.M) {
	os.Setenv("dbpass", "123456")

	pwd := os.Getenv("dbpass")
	if pwd == "" {
		m.Run()
	}

	connstr := "postgres://postgres:" + pwd + "@localhost:5432/task_30_8_1"
	var err error
	s, err = New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestStorage_Tasks(t *testing.T) {
	data, err := s.Tasks(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = s.Tasks(1, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestStorage_NewTask(t *testing.T) {
	task := Task{
		Title: "unit test task title",
		Content: "unit test task content",
	}
	id, err := s.NewTask(task)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создана задача с ИД:", id)
}

func TestStorage_DeleteTask(t *testing.T) {
	err := s.DeleteTask(2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStorage_NewLabel(t *testing.T) {
	label := Label{
		Name: "unit test label name",
	}
	id, err := s.NewLabel(label)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создана метка с ИД:", id)
}

func TestStorage_AddLabelToTask(t *testing.T) {
	err := s.AddLabelToTask(8,8)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Метка добавилась к задаче")
}

func TestStorage_NewAuthor(t *testing.T) {
	author := Author{
		Name: "unit test user name",
	}
	id, err := s.NewAuthor(author)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создан автор с ИД:", id)
}

func TestStorage_UpdateTask(t *testing.T) {
	task := Task{
		Title: "unit test task title",
		Content: "unit test task content",
	}
	id, err := s.UpdateTask(task, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Изменена задача с ИД:", id)
}

func TestStorage_TasksByLabel(t *testing.T) {
	data, err := s.TasksByLabel(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

