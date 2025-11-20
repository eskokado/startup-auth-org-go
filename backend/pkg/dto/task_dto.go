package dto

type CreateTaskParams struct {
    OrganizationID string `json:"organization_id"`
    Title          string `json:"title"`
    Description    string `json:"description"`
}

type UpdateTaskStatusParams struct {
    TaskID         string `json:"task_id"`
    OrganizationID string `json:"organization_id"`
    Title          string `json:"title"`
    Description    string `json:"description"`
    Status         string `json:"status"`
}