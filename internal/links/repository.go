package links

import "purple/pkg/db"

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
func (repo *LinkRepository) GetAll(hash string) (*Link, error) {}
