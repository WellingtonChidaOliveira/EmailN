package campaign

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByID(id int) (Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	Save(campaign *Campaign)  error
	Update(campaign *Campaign) (Campaign, error)
}