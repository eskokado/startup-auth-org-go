package repository

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
)

type InvitationRepository interface {
    Save(ctx context.Context, inv *entity.Invitation) (*entity.Invitation, error)
    GetByToken(ctx context.Context, token string) (*entity.Invitation, error)
}