package campaign

import "time"

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
