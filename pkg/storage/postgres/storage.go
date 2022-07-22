package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Задача.
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// Метка
type Label struct {
	ID         int
	Name      string
}

// Автор
type Author struct {
	ID         int
	Name      string
}

// Tasks возвращает список задач из БД.
func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) AND
			($2 = 0 OR author_id = $2)
		ORDER BY id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// Tasks возвращает список задач по метке
func (s *Storage) TasksByLabel(labelID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			t.id,
			t.opened,
			t.closed,
			t.author_id,
			t.assigned_id,
			t.title,
			t.content
		FROM tasks as t
		JOIN tasks_labels as tl
			ON tl.label_id = $1 AND tl.task_id = t.id
		ORDER BY id;
	`,
		labelID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// NewTask создаёт новую задачу и возвращает её id
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content, author_id)
		VALUES ($1, $2, $3) RETURNING id;
		`,
		t.Title,
		t.Content,
		t.AuthorID,
	).Scan(&id)
	return id, err
}

// UpdateTask обновляет задачу и возвращает её id
func (s *Storage) UpdateTask(t Task, task_id int) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		UPDATE tasks SET title=$1, content=$2, author_id=$3, opened=$4, closed=$5, assigned_id=$6
		WHERE id = $7
			RETURNING id;
		`,
		t.Title,
		t.Content,
		t.AuthorID,
		t.Opened,
		t.Closed,
		t.AssignedID,
		task_id,
	).Scan(&id)
	return id, err
}

// DeleteTask удаляет по её id
func (s *Storage) DeleteTask(id int) (error) {
	_,err := s.db.Exec(context.Background(), `
		DELETE FROM tasks WHERE id=$1;
		`,
		id,
	)
	return err
}

// NewAuthor создаёт нового автора и возвращает её id
func (s *Storage) NewAuthor(a Author) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO users (name)
		VALUES ($1) RETURNING id;
		`,
		a.Name,
	).Scan(&id)
	return id, err
}

// NewLabel создаёт новую метку и возвращает её id
func (s *Storage) NewLabel(l Label) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO labels (name)
		VALUES ($1) RETURNING id;
		`,
		l.Name,
	).Scan(&id)
	return id, err
}

// AddLabelToTask добавляет метку к задачи
func (s *Storage) AddLabelToTask(label_id, task_id int) (error) {
	_,err := s.db.Exec(context.Background(), `
		INSERT INTO tasks_labels (task_id, label_id)
		VALUES ($1, $2);
		`,
		task_id,
		label_id,
	)
	return err
}