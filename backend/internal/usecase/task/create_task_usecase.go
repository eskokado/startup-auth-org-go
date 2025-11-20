package usecase

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/repository"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

type CreateTaskUsecase struct { repo repository.TaskRepository }

func NewCreateTaskUsecase(repo repository.TaskRepository) *CreateTaskUsecase { return &CreateTaskUsecase{repo: repo} }

func (u *CreateTaskUsecase) Execute(ctx context.Context, orgID vo.ID, title string, description string) (*entity.Task, error) {
    t, err := entity.NewTask(orgID, title, description)
    if err != nil { return nil, err }
    return u.repo.Save(ctx, t)
}