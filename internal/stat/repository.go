package stat

import (
	"go-demo-6/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatrepository(database *db.Db) *StatRepository {
	return &StatRepository{
		Db: database,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	var currentdate datatypes.Date = datatypes.Date(time.Now())

	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentdate)

	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Date:   currentdate,
			Clicks: 1,
		})

	} else {
		stat.Clicks++
		repo.Db.Save(&stat)

	}
}

func (repo *StatRepository) GetStats(by string, from, to time.Time) ([]GetStatResponse, error) {
	var stats []GetStatResponse
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}

	result := repo.Db.
		Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)
	if result.Error != nil {
		return nil, result.Error
	}

	return stats, nil
}
