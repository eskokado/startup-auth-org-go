package repository

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

type TaskRepository interface {
    Save(ctx context.Context, t *entity.Task) (*entity.Task, error)
    Update(ctx context.Context, t *entity.Task) (*entity.Task, error)
    Delete(ctx context.Context, id vo.ID) error
    ListByOrganization(ctx context.Context, orgID vo.ID) ([]*entity.Task, error)
}