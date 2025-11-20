package usecase

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/repository"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

type DeleteTaskUsecase struct { repo repository.TaskRepository }

func NewDeleteTaskUsecase(repo repository.TaskRepository) *DeleteTaskUsecase { return &DeleteTaskUsecase{repo: repo} }

func (u *DeleteTaskUsecase) Execute(ctx context.Context, id vo.ID) error { return u.repo.Delete(ctx, id) }