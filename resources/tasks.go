package resources

type Task struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status int16  `json:"status"`
}

const ListTasksQuery = `SELECT id, task, status FROM todos;`
const AddTaskQuery = `INSERT INTO todos (task, status) VALUES (?, ?);`
const RemoveTaskQuery = `DELETE FROM todos WHERE id = ?;`
const UpdateTaskQuery = `UPDATE todos SET status = ? WHERE id = ?;`
const GetTaskQuery = `SELECT * FROM todos WHERE id = ?;`
const GetTaskCountQuery = `SELECT COUNT(*) FROM todos;`
const GetTaskByIDQuery = `SELECT * FROM todos WHERE id = ?;`
const GetTaskByStatusQuery = `SELECT * FROM todos WHERE status = ?;`
