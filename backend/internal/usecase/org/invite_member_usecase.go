package usecase

import (
    "context"
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/repository"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "github.com/eskokado/startup-auth-go/backend/pkg/msgerror"
)

// InviteMemberUsecase envia convite por e-mail para entrar na organização
type InviteMemberUsecase struct {
    invRepo      repository.InvitationRepository
    memberRepo   repository.MembershipRepository
    orgRepo      repository.OrganizationRepository
    inviteTTL    time.Duration
    baseURL      string
}

func NewInviteMemberUsecase(invRepo repository.InvitationRepository, memberRepo repository.MembershipRepository, orgRepo repository.OrganizationRepository, ttl time.Duration, baseURL string) *InviteMemberUsecase {
    return &InviteMemberUsecase{invRepo: invRepo, memberRepo: memberRepo, orgRepo: orgRepo, inviteTTL: ttl, baseURL: baseURL}
}

func (u *InviteMemberUsecase) Execute(ctx context.Context, orgID vo.ID, inviterID vo.ID, emailStr string) (*entity.Invitation, error) {
    email, err := vo.NewEmail(emailStr)
    if err != nil { return nil, err }
    // Verifica plano da organização
    org, err := u.orgRepo.GetByID(ctx, orgID)
    if err != nil { return nil, err }
    if org.Plan.String() == vo.PlanPersonal { return nil, msgerror.AnErrNotAllowed }
    // Evita convidar quem já é membro
    already, err := u.memberRepo.Exists(ctx, orgID, inviterID)
    if err != nil { return nil, err }
    _ = already

    inv, err := entity.NewInvitation(orgID, email, inviterID, u.inviteTTL)
    if err != nil { return nil, err }
    saved, err := u.invRepo.Save(ctx, inv)
    if err != nil { return nil, err }
    // Retorna token para envio externo (frontend ou serviço de email dedicado)
    _ = u.baseURL
    return saved, nil
}