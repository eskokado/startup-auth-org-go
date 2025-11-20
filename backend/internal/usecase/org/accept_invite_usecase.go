package usecase

import (
	"context"
	"time"

	"github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
	"github.com/eskokado/startup-auth-go/backend/pkg/domain/repository"
	"github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
	"github.com/eskokado/startup-auth-go/backend/pkg/msgerror"
)

// AcceptInviteUsecase aceita um convite e cria o vínculo
type AcceptInviteUsecase struct {
	invRepo    repository.InvitationRepository
	memberRepo repository.MembershipRepository
}

func NewAcceptInviteUsecase(invRepo repository.InvitationRepository, memberRepo repository.MembershipRepository) *AcceptInviteUsecase {
	return &AcceptInviteUsecase{invRepo: invRepo, memberRepo: memberRepo}
}

func (u *AcceptInviteUsecase) Execute(ctx context.Context, token string, userID vo.ID) error {
	inv, err := u.invRepo.GetByToken(ctx, token)
	if err != nil {
		return err
	}
	if time.Now().After(inv.ExpiresAt) {
		return msgerror.AnErrExpiredToken
	}
	role, _ := vo.NewRole(vo.RoleMember)
	_, err = u.memberRepo.Save(ctx, entity.NewMembership(inv.OrganizationID, userID, role))
	if err != nil {
		return err
	}
	now := time.Now()
	inv.Accept(now)
	_, _ = u.invRepo.Save(ctx, inv)
	return nil
}
