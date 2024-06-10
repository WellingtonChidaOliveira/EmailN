package campaign

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID          string
	Name        string
	CreatedOn   time.Time
	ModifiedOn  time.Time
	Content     string
	Recipients  []Contact
	Template    string
	IsActivated bool
}

func NewCampaign(name, content, template string, recipients []Contact) *Campaign {

	contacts := make([]Contact, len(recipients))
	for i, r := range recipients {
		contacts[i].Email = r.Email
	}

	return &Campaign{
		ID:          uuid.New().String(),
		Name:        name,
		CreatedOn:   time.Now(),
		ModifiedOn:  time.Now(),
		Content:     content,
		Recipients:  contacts,
		Template:    template,
		IsActivated: true,
	}
}