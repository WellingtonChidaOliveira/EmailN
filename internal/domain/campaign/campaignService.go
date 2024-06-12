package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)

type Service struct {
	Repository Repository
}

func (s* Service) GetCampaigns() ([]Campaign, error) {
	return s.Repository.FindAll()
}

func (s* Service) GetCampaignByID(id int) (Campaign, error) {
	return s.Repository.FindByID(id)
}

func (s* Service) GetCampaignsByUserID(userID int) ([]Campaign, error) {
	return s.Repository.FindByUserID(userID)
}

func (s* Service) CreateCampaign(newCampaign contract.NewCampaignDto) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, "",newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil{
		return "", internalerrors.ErrDataBase
	}
	return campaign.ID, nil
}

func (s* Service) UpdateCampaign(campaign *Campaign) (Campaign, error) {
	return s.Repository.Update(campaign)
}