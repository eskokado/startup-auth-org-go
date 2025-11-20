package repository

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

type MembershipRepository interface {
    Save(ctx context.Context, m *entity.Membership) (*entity.Membership, error)
    Exists(ctx context.Context, orgID vo.ID, userID vo.ID) (bool, error)
}