package contract

// NewCampaignDto is a data transfer object for creating a new campaign.
type NewCampaignDto struct {
	UserID    int    `json:"user_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
	Emails []string `json:"emails" binding:"required"`
}