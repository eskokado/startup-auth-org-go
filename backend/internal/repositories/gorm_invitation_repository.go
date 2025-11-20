package repository

import (
    "context"
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "gorm.io/gorm"
)

type GormInvitation struct {
    ID             string    `gorm:"primaryKey;type:varchar(36)"`
    OrganizationID string    `gorm:"type:varchar(36);index;not null"`
    Email          string    `gorm:"type:varchar(255);index;not null"`
    Token          string    `gorm:"type:varchar(255);uniqueIndex;not null"`
    ExpiresAt      time.Time `gorm:"type:datetime"`
    InviterID      string    `gorm:"type:varchar(36);not null"`
    AcceptedAt     *time.Time
}

type GormInvitationRepository struct{ db *gorm.DB }

func NewGormInvitationRepository(db *gorm.DB) *GormInvitationRepository { return &GormInvitationRepository{db: db} }

func (r *GormInvitationRepository) toDBModel(e *entity.Invitation) *GormInvitation {
    return &GormInvitation{ID: e.ID.String(), OrganizationID: e.OrganizationID.String(), Email: e.Email.String(), Token: e.Token, ExpiresAt: e.ExpiresAt, InviterID: e.InviterID.String(), AcceptedAt: e.AcceptedAt}
}

func (r *GormInvitationRepository) fromDBModel(m *GormInvitation) *entity.Invitation {
    // Simplificação: ignorando erros de VO para foco na persistência
    inv := &entity.Invitation{}
    inv.ID, _ = vo.ParseID(m.ID)
    inv.OrganizationID, _ = vo.ParseID(m.OrganizationID)
    inv.Email, _ = vo.NewEmail(m.Email)
    inv.Token = m.Token
    inv.ExpiresAt = m.ExpiresAt
    inv.InviterID, _ = vo.ParseID(m.InviterID)
    inv.AcceptedAt = m.AcceptedAt
    return inv
}

func (r *GormInvitationRepository) Save(ctx context.Context, e *entity.Invitation) (*entity.Invitation, error) {
    model := r.toDBModel(e)
    if err := r.db.WithContext(ctx).Save(model).Error; err != nil { return nil, err }
    return r.fromDBModel(model), nil
}

func (r *GormInvitationRepository) GetByToken(ctx context.Context, token string) (*entity.Invitation, error) {
    var m GormInvitation
    if err := r.db.WithContext(ctx).Where("token = ?", token).First(&m).Error; err != nil { return nil, err }
    return r.fromDBModel(&m), nil
}