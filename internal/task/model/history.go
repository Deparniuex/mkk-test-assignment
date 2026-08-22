package model

type TaskHistory struct {
	ID        uint `db:"id"`
	TaskID    uint `db:"task_id"`
	ChangedBy uint `db:"changed_by"`
	Changes   Task `db:"changes"`
}
