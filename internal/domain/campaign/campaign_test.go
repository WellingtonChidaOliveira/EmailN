package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCampaign(t *testing.T) {

	tests := []struct {
		name     string
		setup    func() (*Campaign, error)
		validate func(campaign *Campaign, err error)
	}{
		{
			name: "NewCampaign",
			setup: func() (*Campaign, error) {
				return GenerateCampaign()
			},
			validate: func(campaign *Campaign, err error) {
				require := require.New(t)
				assert := assert.New(t)

				require.NotNil(campaign, "NewCampaign should return a non-nil campaign")
				assert.Equal(campaign.Name, name)
				assert.Greater(len(campaign.Recipients), 0)
				assert.Equal(len(campaign.Recipients), len(recipients))
				assert.Equal(campaign.Recipients[0].Email, recipients[0])
			},
		},
		{
			name: "GenerateId",
			setup: func() (*Campaign, error) {
				return GenerateCampaign()
			},
			validate: func(campaign *Campaign, err error) {
				require := require.New(t)
				assert := assert.New(t)

				require.NotNil(campaign, "NewCampaign should return a non-nil campaign")
				assert.NotEmpty(campaign.ID)
			},
		},
		{
			name: "GenerateCreatedOn",
			setup: func() (*Campaign, error) {
				return GenerateCampaign()
			},
			validate: func(campaign *Campaign, err error) {
				require := require.New(t)
				assert := assert.New(t)

				require.NotNil(campaign, "NewCampaign should return a non-nil campaign")
				assert.NotEmpty(campaign.CreatedOn)
				assert.Greater(campaign.CreatedOn.Unix(), int64(0))
			},
		},
		{
			name: "ValidateName",
			setup: func() (*Campaign, error) {
				return NewCampaign("", content, template, recipients)
			},
			validate: func(campaign *Campaign, err error) {
				assert := assert.New(t)

				assert.Equal("Name is required", err.Error())
			},
		},
		{
			name: "ValidateRecipients",
			setup: func() (*Campaign, error) {
				return NewCampaign(name, content, template, nil)
			},
			validate: func(campaign *Campaign, err error) {
				require := require.New(t)
				assert := assert.New(t)

				require.NotNil(err)
				assert.Equal(ErrRecipientsAreRequired, err.Error())
			},
		},
		{
			name: "ValidateEmail",
			setup: func() (*Campaign, error) {
				return NewCampaign(name, content, template, []string{"w"})
			},
			validate: func(campaign *Campaign, err error) {
				require := require.New(t)
				assert := assert.New(t)

				require.NotNil(err)
				assert.Equal(ErrInvalidEmail, err.Error())
			},
		},
		{
			name: "ValidateEmailIsRequired",
			setup: func() (*Campaign, error) {
				return NewCampaign(name, content, template, []string{""})
			},
			validate: func(campaign *Campaign, err error) {
				require := require.New(t)
				assert := assert.New(t)

				require.NotNil(err)
				assert.Equal(ErrInvalidEmail, err.Error())
			},
		},
		{
			name: "ValidateContent",
			setup: func() (*Campaign, error) {
				return NewCampaign(name, "  ", template, recipients)
			},
			validate: func(campaign *Campaign, err error) {
				assert := assert.New(t)

				assert.Equal("Content is too short", err.Error())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			campaign, err := tt.setup()
			tt.validate(campaign, err)
		})
	}
}

var (
	name       = "Test Campaign"
	content    = "Test Content"
	template   = "Test Template"
	recipients = []string{"teste@test.com"}
)

func GenerateCampaign() (*Campaign, error) {
	return NewCampaign(name, content, template, recipients)
}
