package repository

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

type OrganizationRepository interface {
    Save(ctx context.Context, org *entity.Organization) (*entity.Organization, error)
    GetByID(ctx context.Context, id vo.ID) (*entity.Organization, error)
    GetByOwnerID(ctx context.Context, ownerID vo.ID) (*entity.Organization, error)
}