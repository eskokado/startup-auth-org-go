package vo

import (
    "encoding/json"
    "github.com/eskokado/startup-auth-go/backend/pkg/msgerror"
)

// TaskStatus representa o estado da tarefa
type TaskStatus struct {
    value string
}

const (
    TaskTodo       = "TODO"
    TaskInProgress = "IN_PROGRESS"
    TaskDone       = "DONE"
)

func NewTaskStatus(value string) (TaskStatus, error) {
    switch value {
    case TaskTodo, TaskInProgress, TaskDone:
        return TaskStatus{value: value}, nil
    default:
        return TaskStatus{}, msgerror.AnErrInvalidStatus
    }
}

func (s TaskStatus) String() string { return s.value }

func (s TaskStatus) IsEmpty() bool { return s.value == "" }

func (s TaskStatus) MarshalJSON() ([]byte, error) { return json.Marshal(s.value) }