package models

// Guerateam model
type Guerateam struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OwnerID     string   `json:"ownerId"`
	MembersID   []string `json:"membersId"`
}
