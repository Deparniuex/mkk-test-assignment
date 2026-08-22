package model

type TaskFilter struct {
	AssigneeID uint
	Status     *Status
	TeamID     uint
}
