package campaign

import (
	"emailn/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

var (
	mockRepo = new(MockRepository)
	service = Service{Repository: mockRepo}

)

func Test_Create(t *testing.T){
	t.Run("CreateCampaign", func(t *testing.T){
		//arrange
		assert := assert.New(t)
		mockRepo.On("Save", mock.Anything).Return(nil)
		newCampaign := generateNewCampaign()

		// Act
		id, err := service.CreateCampaign(newCampaign)

		// Assert
		assert.Nil(err)
		assert.NotEmpty(id)
	})

	t.Run("CreateCampaignSaveInDatabase", func(t *testing.T){
		//arrange
		newCampaign := generateNewCampaign()
		mockRepo.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
			return campaign.Name == newCampaign.Name && campaign.Content == newCampaign.Content
		})).Return(nil)

		// Act
		service.CreateCampaign(newCampaign)

		mockRepo.AssertExpectations(t)
	})

	t.Run("ValidateErrorWhenCreate", func(t *testing.T) {
		assert := assert.New(t)
		newCampaign := generateNewCampaign()
		newCampaign.Name = ""
		
		_, err := service.CreateCampaign(newCampaign)

		assert.NotNil(err)
		assert.Equal("name is required", err.Error())
	})

	t.Run("ValidateErrorWhenSaveDb", func(t *testing.T){
		assert := assert.New(t)
		mockRepo := new(MockRepository)
		mockRepo.On("Save", mock.Anything).Return(errors.New("Error to save on db"))
		service := Service{Repository: mockRepo}

		newCampaign := generateNewCampaign()
		_, err := service.CreateCampaign(newCampaign)

		if err == nil {
			t.Errorf("Expected err nut I got nil")
		}
		assert.NotNil(err)
		assert.Equal("Error to save on db", err.Error())
		
		
	})
}

func generateNewCampaign() contract.NewCampaignDto {
	return contract.NewCampaignDto{
		Name : name,
		Content: content,
		Emails: []string{recipients[0]},
	}
}


func (m *MockRepository) FindAll() ([]Campaign, error) {
	args := m.Called()
	return args.Get(0).([]Campaign), args.Error(1)
}

func (m *MockRepository) FindByID(id int) (Campaign, error) {
	args := m.Called(id)
	return args.Get(0).(Campaign), args.Error(1)
}

func (m *MockRepository) FindByUserID(userID int) ([]Campaign, error) {
	args := m.Called(userID)
	return args.Get(0).([]Campaign), args.Error(1)
}

func (m *MockRepository) Save(campaign *Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *MockRepository) Update(campaign *Campaign) (Campaign, error) {
	args := m.Called(campaign)
	return args.Get(0).(Campaign), args.Error(1)
}