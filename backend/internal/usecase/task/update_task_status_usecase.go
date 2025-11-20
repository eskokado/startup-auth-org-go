package usecase

import (
    "context"
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/repository"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

type UpdateTaskStatusUsecase struct { repo repository.TaskRepository }

func NewUpdateTaskStatusUsecase(repo repository.TaskRepository) *UpdateTaskStatusUsecase { return &UpdateTaskStatusUsecase{repo: repo} }

func (u *UpdateTaskStatusUsecase) Execute(ctx context.Context, t *entity.Task, status string) (*entity.Task, error) {
    s, err := vo.NewTaskStatus(status)
    if err != nil { return nil, err }
    t.Status = s
    t.UpdatedAt = time.Now()
    return u.repo.Update(ctx, t)
}