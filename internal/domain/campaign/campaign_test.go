package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCampaign(t *testing.T) {
	t.Run("NewCampaign", func(t *testing.T){
		assert := assert.New(t)
		require := require.New(t)

		//act
		campaign, _  := GenerateCampaign()

		//assert
		require.NotNil(campaign, "NewCampaign should return a non-nil campaign")
		assert.Equal(campaign.Name, name)
		assert.Greater(len(campaign.Recipients), 0)
		assert.Equal(len(campaign.Recipients), len(recipients))
		assert.Equal(campaign.Recipients[0].Email, recipients[0].Email)
	})

	t.Run("GenerateId", func(t *testing.T){
		assert := assert.New(t)
		require := require.New(t)
		
		//act
		campaign, _ := GenerateCampaign()

		//assert
		require.NotNil(campaign, "NewCampaign should return a non-nil campaign")
		assert.NotEmpty(campaign.ID)
	})

	t.Run("GenerateCreatedOn", func(t *testing.T){
		assert := assert.New(t)
		require := require.New(t)
		
		//act
		campaign, _ := GenerateCampaign()
		

		//assert
		require.NotNil(campaign, "NewCampaign should return a non-nil campaign")
		assert.NotEmpty(campaign.CreatedOn)
		assert.Greater(campaign.CreatedOn.Unix(), int64(0))
	})

	t.Run("ValidateName", func(t *testing.T){
		assert := assert.New(t)
		//require := require.New(t)

		//act
		_, err := NewCampaign("  ", content, template, recipients)
		
		//assert
		assert.Equal("name is required", err.Error())
	})

	t.Run("ValidateRecipients", func(t *testing.T){
		assert := assert.New(t)
		require := require.New(t)

		//act
		_, err := NewCampaign(name, content, template, []Contact{})
		
		//assert
		require.NotNil(err)
		assert.Equal(ErrRecipientsAreRequired, err.Error())
	})

	t.Run("ValidateEmail", func(t *testing.T){
		assert := assert.New(t)
		require := require.New(t)

		//act
		_, err := NewCampaign(name, content, template, []Contact{
			{
				Email: "teste",
			},
		})
		
		//assert
		require.NotNil(err)
		assert.Equal(ErrInvalidEmail, err.Error())
	})

	t.Run("ValidateEmailIsRequired", func(t *testing.T){
		assert := assert.New(t)
		require := require.New(t)

		//act
		_, err := NewCampaign(name, content, template, []Contact{
			{
				Email: "",
			},
		})
		
		//assert
		require.NotNil(err)
		assert.Equal(ErrEmailIsRequired, err.Error())
	})

	t.Run("ValidateContent", func(t *testing.T){
		assert := assert.New(t)
		//require := require.New(t)

		//act
		_, err := NewCampaign(name, "  ", template, recipients)
		
		//assert
		assert.Equal("content is required", err.Error())
	})
		
}

var(
	name = "Test Campaign"
	content = "Test Content"
	template = "Test Template"
	recipients = []Contact{
		{
			Email: "teste@test.com",
		},
	}
)

func GenerateCampaign() (*Campaign, error) {
	return NewCampaign(name, content, template, recipients)
}