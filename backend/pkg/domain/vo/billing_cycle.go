package vo

import "github.com/eskokado/startup-auth-go/backend/pkg/msgerror"

type BillingCycle struct{ value string }

const (
    CycleMonthly    = "MONTHLY"
    CycleSemiannual = "SEMIANNUAL"
    CycleAnnual     = "ANNUAL"
)

func NewBillingCycle(v string) (BillingCycle, error) {
    switch v {
    case CycleMonthly, CycleSemiannual, CycleAnnual:
        return BillingCycle{value: v}, nil
    default:
        return BillingCycle{}, msgerror.AnErrInvalidCycle
    }
}

func (c BillingCycle) String() string { return c.value }
func (c BillingCycle) IsEmpty() bool { return c.value == "" }