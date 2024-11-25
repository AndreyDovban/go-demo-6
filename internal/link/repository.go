package link

import (
	"go-demo-6/pkg/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	*db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Db: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Db.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Db.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Db.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	var link Link
	result := repo.Db.Delete(&link, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *LinkRepository) GetById(id uint) error {
	var link Link
	result := repo.Db.First(&link, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *LinkRepository) Count() (int64, error) {
	var count int64
	result := repo.Db.
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil

}

func (repo *LinkRepository) GetAll(limit, offset int) ([]Link, error) {
	var links []Link
	query := repo.Db.
		Table("links").
		Where("deleted_at is null").
		Session(&gorm.Session{})

	result := query.
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}
