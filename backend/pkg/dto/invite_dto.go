package dto

// InviteParams para convidar membro
type InviteParams struct {
    OrganizationID string `json:"organization_id"`
    InviterID      string `json:"inviter_id"`
    Email          string `json:"email"`
}

type AcceptInviteParams struct {
    Token  string `json:"token"`
    UserID string `json:"user_id"`
}