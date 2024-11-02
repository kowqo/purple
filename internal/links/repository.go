package links

import (
	"gorm.io/gorm/clause"
	"purple/pkg/db"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	res := repo.Database.DB.Create(link)

	if res.Error != nil {
		return nil, res.Error
	}

	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	link := &Link{}
	res := repo.Database.DB.First(link, "hash = ?", hash)
	if res.Error != nil {
		return nil, res.Error
	}

	return link, nil

}
func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	res := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if res.Error != nil {
		return nil, res.Error
	}

	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {

	res := repo.Database.DB.Delete(&Link{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo *LinkRepository) FindById(id uint) (*Link, error) {
	var link Link
	res := repo.Database.DB.First(&link, id)

	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}
