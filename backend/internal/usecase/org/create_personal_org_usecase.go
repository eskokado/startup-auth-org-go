package usecase

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/repository"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

// CreatePersonalOrgUsecase cria a organização padrão "Personal" para o usuário
type CreatePersonalOrgUsecase struct {
    orgRepo       repository.OrganizationRepository
    memberRepo    repository.MembershipRepository
}

func NewCreatePersonalOrgUsecase(orgRepo repository.OrganizationRepository, memberRepo repository.MembershipRepository) *CreatePersonalOrgUsecase {
    return &CreatePersonalOrgUsecase{orgRepo: orgRepo, memberRepo: memberRepo}
}

func (u *CreatePersonalOrgUsecase) Execute(ctx context.Context, userID vo.ID) error {
    name, _ := vo.NewOrganizationName("Personal")
    org := entity.NewOrganization(name, userID)
    saved, err := u.orgRepo.Save(ctx, org)
    if err != nil { return err }
    role, _ := vo.NewRole(vo.RoleOwner)
    _, err = u.memberRepo.Save(ctx, entity.NewMembership(saved.ID, userID, role))
    return err
}