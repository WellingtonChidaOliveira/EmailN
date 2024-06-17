package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

var (
	ErrNameIsRequired        = "name is required"
	ErrContentIsRequired     = "content is required"
	ErrRecipientsAreRequired = "Recipients is too short"
	ErrInvalidEmail          = "Email is invalid"
	ErrEmailIsRequired       = "email is required"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID          string    `validate:"required"`
	Name        string    `validate:"required,min=5,max=100"`
	CreatedOn   time.Time `validate:"required"`
	ModifiedOn  time.Time `validate:"required"`
	Content     string    `validate:"min=5"`
	Recipients  []Contact `validate:"min=1,dive"`
	Template    string
	IsActivated bool
}

func NewCampaign(name, content, template string, recipients []string) (*Campaign, error) {

	contacts := make([]Contact, 0, len(recipients))
	for _, recipient := range recipients {
		contacts = append(contacts, Contact{Email: recipient})
	}

	campaign := &Campaign{
		ID:          xid.New().String(),
		Name:        name,
		CreatedOn:   time.Now(),
		ModifiedOn:  time.Now(),
		Content:     content,
		Recipients:  contacts,
		Template:    template,
		IsActivated: true,
	}

	err := internalerrors.ValidatorStruct(campaign)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
