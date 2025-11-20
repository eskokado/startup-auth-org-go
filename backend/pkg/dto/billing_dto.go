package dto

type CheckoutParams struct {
    OrganizationID string `json:"organization_id"`
    Plan           string `json:"plan"`
    Cycle          string `json:"cycle"`
    SuccessURL     string `json:"success_url"`
    CancelURL      string `json:"cancel_url"`
}