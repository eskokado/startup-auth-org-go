package entity

import (
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

// Organization representa uma organização SaaS
type Organization struct {
    ID        vo.ID
    Name      vo.OrganizationName
    OwnerID   vo.ID
    Plan      vo.PlanType
    CreatedAt time.Time
}

func NewOrganization(name vo.OrganizationName, ownerID vo.ID) *Organization {
    return &Organization{
        ID:        vo.NewID(),
        Name:      name,
        OwnerID:   ownerID,
        Plan:      func() vo.PlanType { p, _ := vo.NewPlanType(vo.PlanPersonal); return p }(),
        CreatedAt: time.Now(),
    }
}