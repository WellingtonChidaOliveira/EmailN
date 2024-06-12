package campaign

import (
	"errors"
	"strings"
	"time"

	"github.com/rs/xid"
)

var (
	ErrNameIsRequired = "name is required"
	ErrContentIsRequired = "content is required"
	ErrRecipientsAreRequired = "recipients are required"
	ErrInvalidEmail = "invalid email"
	ErrEmailIsRequired = "email is required"
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

func NewCampaign(name, content, template string, recipients []string) (*Campaign, error) {

	if len(recipients) == 0 {
		return nil, errors.New(ErrRecipientsAreRequired)
	}

	contacts := make([]Contact, len(recipients))
	for i, r := range recipients {
		contacts[i].Email = r
	}

	if strings.TrimSpace(name) == "" {
		return nil, errors.New(ErrNameIsRequired)
	} else if strings.TrimSpace(content) == "" {
		return nil, errors.New(ErrContentIsRequired)
	}

	for _, c := range contacts {
		if strings.TrimSpace(c.Email) == "" {
			return nil, errors.New(ErrEmailIsRequired)
		} else if !strings.Contains(c.Email, "@") {
			return nil, errors.New(ErrInvalidEmail)
		}
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