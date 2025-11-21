package vo

import (
    "encoding/json"
    "fmt"
    "strings"

    "github.com/eskokado/startup-auth-go/backend/pkg/msgerror"
)

// OrganizationName representa o nome de uma organização
type OrganizationName struct {
    value string
}

func NewOrganizationName(value string) (OrganizationName, error) {
    trimmed := strings.TrimSpace(value)
    if trimmed == "" {
        return OrganizationName{}, msgerror.AnErrEmptyName
    }
    if len(trimmed) < 3 {
        return OrganizationName{}, fmt.Errorf("%w: mínimo 3 caracteres", msgerror.AnErrNameTooShort)
    }
    if len(trimmed) > 100 {
        return OrganizationName{}, fmt.Errorf("%w: máximo 100 caracteres", msgerror.AnErrNameTooLong)
    }
    return OrganizationName{value: trimmed}, nil
}

func (n OrganizationName) String() string {
    return n.value
}

func (n OrganizationName) IsEmpty() bool {
    return n.value == ""
}

func (n OrganizationName) MarshalJSON() ([]byte, error) { return json.Marshal(n.value) }