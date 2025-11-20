package vo

import "github.com/eskokado/startup-auth-go/backend/pkg/msgerror"

// Role representa o papel do membro na organização
type Role struct {
    value string
}

const (
    RoleOwner  = "OWNER"
    RoleAdmin  = "ADMIN"
    RoleMember = "MEMBER"
)

func NewRole(value string) (Role, error) {
    switch value {
    case RoleOwner, RoleAdmin, RoleMember:
        return Role{value: value}, nil
    default:
        return Role{}, msgerror.AnErrInvalidRole
    }
}

func (r Role) String() string { return r.value }

func (r Role) IsEmpty() bool { return r.value == "" }