package usecase

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/repository"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

type ListTasksUsecase struct { repo repository.TaskRepository }

func NewListTasksUsecase(repo repository.TaskRepository) *ListTasksUsecase { return &ListTasksUsecase{repo: repo} }

func (u *ListTasksUsecase) Execute(ctx context.Context, orgID vo.ID) ([]*entity.Task, error) {
    return u.repo.ListByOrganization(ctx, orgID)
}