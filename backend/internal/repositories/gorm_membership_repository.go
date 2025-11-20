package repository

import (
    "context"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "gorm.io/gorm"
)

type GormMembership struct {
    ID             string `gorm:"primaryKey;type:varchar(36)"`
    OrganizationID string `gorm:"type:varchar(36);index;not null"`
    UserID         string `gorm:"type:varchar(36);index;not null"`
    Role           string `gorm:"type:varchar(20);not null"`
}

type GormMembershipRepository struct{ db *gorm.DB }

func NewGormMembershipRepository(db *gorm.DB) *GormMembershipRepository { return &GormMembershipRepository{db: db} }

func (r *GormMembershipRepository) toDBModel(e *entity.Membership) *GormMembership {
    return &GormMembership{ID: e.ID.String(), OrganizationID: e.OrganizationID.String(), UserID: e.UserID.String(), Role: e.Role.String()}
}

func (r *GormMembershipRepository) fromDBModel(m *GormMembership) (*entity.Membership, error) {
    id, _ := vo.ParseID(m.ID)
    orgID, _ := vo.ParseID(m.OrganizationID)
    userID, _ := vo.ParseID(m.UserID)
    role, _ := vo.NewRole(m.Role)
    return &entity.Membership{ID: id, OrganizationID: orgID, UserID: userID, Role: role}, nil
}

func (r *GormMembershipRepository) Save(ctx context.Context, e *entity.Membership) (*entity.Membership, error) {
    model := r.toDBModel(e)
    if err := r.db.WithContext(ctx).Save(model).Error; err != nil { return nil, err }
    return r.fromDBModel(model)
}

func (r *GormMembershipRepository) Exists(ctx context.Context, orgID vo.ID, userID vo.ID) (bool, error) {
    var count int64
    if err := r.db.WithContext(ctx).Model(&GormMembership{}).Where("organization_id = ? AND user_id = ?", orgID.String(), userID.String()).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}