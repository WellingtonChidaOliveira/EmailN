package campaign

import (
	"time"

	"github.com/rs/xid"
)

var (
	ErrNameIsRequired        = "name is required"
	ErrContentIsRequired     = "content is required"
	ErrRecipientsAreRequired = "recipients are required"
	ErrInvalidEmail          = "invalid email"
	ErrEmailIsRequired       = "email is required"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID          string    `validate:"required"`
	Name        string    `validate:"min=5,max=100"`
	CreatedOn   time.Time `validate:"required"`
	ModifiedOn  time.Time `validate:"required"`
	Content     string    `validate:"min=5"`
	Recipients  []Contact `validate:"min=1"`
	Template    string
	IsActivated bool
}

func NewCampaign(name, content, template string, recipients []string) (*Campaign, error) {

	contacts := make([]Contact, len(recipients))
	for _, recipient := range recipients {
		contacts = append(contacts, Contact{Email: recipient})
	}

	return &Campaign{
		ID:          xid.New().String(),
		Name:        name,
		CreatedOn:   time.Now(),
		ModifiedOn:  time.Now(),
		Content:     content,
		Recipients:  contacts,
		Template:    template,
		IsActivated: true,
	}, nil
}
