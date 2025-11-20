package entity

import (
    "crypto/rand"
    "encoding/base64"
    "time"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
)

// Invitation para convidar membro por e-mail
type Invitation struct {
    ID             vo.ID
    OrganizationID vo.ID
    Email          vo.Email
    Token          string
    ExpiresAt      time.Time
    InviterID      vo.ID
    AcceptedAt     *time.Time
}

func NewInvitation(orgID vo.ID, email vo.Email, inviterID vo.ID, ttl time.Duration) (*Invitation, error) {
    token, err := generateToken(32)
    if err != nil {
        return nil, err
    }
    expires := time.Now().Add(ttl)
    return &Invitation{
        ID:             vo.NewID(),
        OrganizationID: orgID,
        Email:          email,
        Token:          token,
        ExpiresAt:      expires,
        InviterID:      inviterID,
    }, nil
}

func (i *Invitation) Accept(now time.Time) {
    i.AcceptedAt = &now
}

func generateToken(n int) (string, error) {
    b := make([]byte, n)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}