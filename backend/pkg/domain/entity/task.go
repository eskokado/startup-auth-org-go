package entity

import (
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

// Task simples por organização
type Task struct {
    ID             vo.ID        `json:"id"`
    OrganizationID vo.ID        `json:"organization_id"`
    Title          string       `json:"title"`
    Description    string       `json:"description"`
    Status         vo.TaskStatus `json:"status"`
    AssigneeID     *vo.ID       `json:"assignee_id"`
    CreatedAt      time.Time    `json:"created_at"`
    UpdatedAt      time.Time    `json:"updated_at"`
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