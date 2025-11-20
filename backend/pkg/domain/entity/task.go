package entity

import (
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

// Task simples por organização
type Task struct {
    ID             vo.ID
    OrganizationID vo.ID
    Title          string
    Description    string
    Status         vo.TaskStatus
    AssigneeID     *vo.ID
    CreatedAt      time.Time
    UpdatedAt      time.Time
}

func NewTask(orgID vo.ID, title string, description string) (*Task, error) {
    status, _ := vo.NewTaskStatus(vo.TaskTodo)
    now := time.Now()
    return &Task{
        ID:             vo.NewID(),
        OrganizationID: orgID,
        Title:          title,
        Description:    description,
        Status:         status,
        CreatedAt:      now,
        UpdatedAt:      now,
    }, nil
}