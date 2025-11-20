package repository

import (
	"context"
	"time"

	"github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
	"github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
	"gorm.io/gorm"
)

type GormOrganization struct {
    ID        string    `gorm:"primaryKey;type:varchar(36)"`
    Name      string    `gorm:"type:varchar(100);not null"`
    OwnerID   string    `gorm:"type:varchar(36);not null"`
    Plan      string    `gorm:"type:varchar(20);not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}

type GormOrganizationRepository struct{ db *gorm.DB }

func NewGormOrganizationRepository(db *gorm.DB) *GormOrganizationRepository {
	return &GormOrganizationRepository{db: db}
}

func (r *GormOrganizationRepository) toDBModel(e *entity.Organization) *GormOrganization {
    return &GormOrganization{ID: e.ID.String(), Name: e.Name.String(), OwnerID: e.OwnerID.String(), Plan: e.Plan.String(), CreatedAt: e.CreatedAt}
}

func (r *GormOrganizationRepository) fromDBModel(m *GormOrganization) (*entity.Organization, error) {
    id, _ := vo.ParseID(m.ID)
    ownerID, _ := vo.ParseID(m.OwnerID)
    name, _ := vo.NewOrganizationName(m.Name)
    plan, _ := vo.NewPlanType(m.Plan)
    return &entity.Organization{ID: id, Name: name, OwnerID: ownerID, Plan: plan, CreatedAt: m.CreatedAt}, nil
}

func (r *GormOrganizationRepository) Save(ctx context.Context, org *entity.Organization) (*entity.Organization, error) {
	model := r.toDBModel(org)
	if err := r.db.WithContext(ctx).Save(model).Error; err != nil {
		return nil, err
	}
	return r.fromDBModel(model)
}

func (r *GormOrganizationRepository) GetByID(ctx context.Context, id vo.ID) (*entity.Organization, error) {
	var m GormOrganization
	if err := r.db.WithContext(ctx).Where("id = ?", id.String()).First(&m).Error; err != nil {
		return nil, err
	}
	return r.fromDBModel(&m)
}

func (r *GormOrganizationRepository) GetByOwnerID(ctx context.Context, ownerID vo.ID) (*entity.Organization, error) {
	var m GormOrganization
	if err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID.String()).First(&m).Error; err != nil {
		return nil, err
	}
	return r.fromDBModel(&m)
}
