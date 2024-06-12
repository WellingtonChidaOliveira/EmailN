package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func Test_Create_Campaign_Db(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(mockRepo *MockRepository, service *Service)
		newCampaign contract.NewCampaignDto
		wantErr  bool
		errText  string
	}{
		{
			name: "CreateCampaign",
			setup: func(mockRepo *MockRepository, service *Service) {
				mockRepo.On("Save", mock.Anything).Return(nil)
			},
			newCampaign: generateNewCampaign(),
			wantErr: false,
		},
		{
			name: "ValidateErrorWhenCreate",
			setup: func(mockRepo *MockRepository, service *Service) {},
			newCampaign: func() contract.NewCampaignDto {
				c := generateNewCampaign()
				c.Name = ""
				return c
			}(),
			wantErr: true,
			errText: "name is required",
		},
		{
			name: "ValidateErrorWhenSaveDb",
			setup: func(mockRepo *MockRepository, service *Service) {
				mockRepo.On("Save", mock.Anything).Return(errors.New("Error to save on db"))
			},
			newCampaign: generateNewCampaign(),
			wantErr: true,
			errText: internalerrors.ErrDataBase.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			service := Service{Repository: mockRepo}
			tt.setup(mockRepo, &service)

			_, err := service.CreateCampaign(tt.newCampaign)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Equal(t, tt.errText, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
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