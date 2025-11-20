package entity

import (
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

// Membership representa vínculo de usuário com organização
type Membership struct {
    ID             vo.ID
    OrganizationID vo.ID
    UserID         vo.ID
    Role           vo.Role
    CreatedAt      time.Time
}

func NewMembership(orgID vo.ID, userID vo.ID, role vo.Role) *Membership {
    return &Membership{
        ID:             vo.NewID(),
        OrganizationID: orgID,
        UserID:         userID,
        Role:           role,
        CreatedAt:      time.Now(),
    }
}