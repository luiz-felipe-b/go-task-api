package tasks

import (
	"database/sql"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) GetAll() ([]Task, error) {
	rows, err := r.DB.Query("SELECT id, title, done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) Create(task Task) (int, error) {
	result, err := r.DB.Exec("INSERT INTO tasks (title, done) VALUES (?, ?)", task.Title, task.Done)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}
