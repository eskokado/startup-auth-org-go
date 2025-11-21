package vo

import (
    "encoding/json"
    "github.com/eskokado/startup-auth-go/backend/pkg/msgerror"
)

type PlanType struct{ value string }

const (
    PlanPersonal     = "PERSONAL"
    PlanOrganization = "ORGANIZATION"
)

func NewPlanType(v string) (PlanType, error) {
    switch v {
    case PlanPersonal, PlanOrganization:
        return PlanType{value: v}, nil
    default:
        return PlanType{}, msgerror.AnErrInvalidPlan
    }
}

func (p PlanType) String() string { return p.value }
func (p PlanType) IsEmpty() bool { return p.value == "" }
func (p PlanType) MarshalJSON() ([]byte, error) { return json.Marshal(p.value) }