package repository

import (
    "context"
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "gorm.io/gorm"
)

type GormTask struct {
    ID             string    `gorm:"primaryKey;type:varchar(36)"`
    OrganizationID string    `gorm:"type:varchar(36);index;not null"`
    Title          string    `gorm:"type:varchar(120);not null"`
    Description    string    `gorm:"type:text"`
    Status         string    `gorm:"type:varchar(20);not null"`
    AssigneeID     *string   `gorm:"type:varchar(36)"`
    CreatedAt      time.Time `gorm:"autoCreateTime"`
    UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

type GormTaskRepository struct{ db *gorm.DB }

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository { return &GormTaskRepository{db: db} }

func (r *GormTaskRepository) toDBModel(e *entity.Task) *GormTask {
    var assignee *string
    if e.AssigneeID != nil { v := e.AssigneeID.String(); assignee = &v }
    return &GormTask{ID: e.ID.String(), OrganizationID: e.OrganizationID.String(), Title: e.Title, Description: e.Description, Status: e.Status.String(), AssigneeID: assignee, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt}
}

func (r *GormTaskRepository) fromDBModel(m *GormTask) (*entity.Task, error) {
    id, _ := vo.ParseID(m.ID)
    orgID, _ := vo.ParseID(m.OrganizationID)
    status, _ := vo.NewTaskStatus(m.Status)
    var assignee *vo.ID
    if m.AssigneeID != nil { v, _ := vo.ParseID(*m.AssigneeID); assignee = &v }
    return &entity.Task{ID: id, OrganizationID: orgID, Title: m.Title, Description: m.Description, Status: status, AssigneeID: assignee, CreatedAt: m.CreatedAt, UpdatedAt: m.UpdatedAt}, nil
}

func (r *GormTaskRepository) Save(ctx context.Context, t *entity.Task) (*entity.Task, error) {
    model := r.toDBModel(t)
    if err := r.db.WithContext(ctx).Save(model).Error; err != nil { return nil, err }
    return r.fromDBModel(model)
}

func (r *GormTaskRepository) Update(ctx context.Context, t *entity.Task) (*entity.Task, error) {
    model := r.toDBModel(t)
    if err := r.db.WithContext(ctx).Save(model).Error; err != nil { return nil, err }
    return r.fromDBModel(model)
}

func (r *GormTaskRepository) Delete(ctx context.Context, id vo.ID) error {
    return r.db.WithContext(ctx).Delete(&GormTask{}, "id = ?", id.String()).Error
}

func (r *GormTaskRepository) ListByOrganization(ctx context.Context, orgID vo.ID) ([]*entity.Task, error) {
    var ms []GormTask
    if err := r.db.WithContext(ctx).Where("organization_id = ?", orgID.String()).Order("updated_at DESC").Find(&ms).Error; err != nil { return nil, err }
    var out []*entity.Task
    for i := range ms {
        t, _ := r.fromDBModel(&ms[i])
        out = append(out, t)
    }
    return out, nil
}