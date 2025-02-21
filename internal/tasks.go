package internal

import (
	"database/sql"

	"github.com/javiersrf/todo-cli/resources"
)

func GetTasks(db *sql.DB) ([]resources.Task, error) {
	rows, err := db.Query(resources.ListTasksQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []resources.Task
	for rows.Next() {
		var task resources.Task
		if err := rows.Scan(&task.ID, &task.Task, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func AddTask(db *sql.DB, task string) error {
	_, err := db.Exec(resources.AddTaskQuery, task, 0)
	return err
}

func MarkTask(db *sql.DB, taskId int, status int) error {
	_, err := db.Exec(resources.UpdateTaskQuery, status, taskId)
	return err
}

func GetTask(db *sql.DB, taskId int) (resources.Task, error) {
	var task resources.Task
	err := db.QueryRow(resources.GetTaskQuery, taskId).Scan(&task.ID, &task.Task, &task.Status)
	return task, err
}
