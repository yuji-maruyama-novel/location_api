package location

import (
	"fmt"
	"gqlgen/graph/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationRepository struct {
	DB *gorm.DB
}

type LocationInterface interface {
	createLocation(model.NewLocation) (*model.Location, error)
	getLocations() ([]*model.Location, error)
	getYears() ([]string, error)
}

func newLocationRepository(db *gorm.DB) *LocationRepository {
	repositry := &LocationRepository{
		DB: db,
	}
	return repositry
}

func (lr *LocationRepository) createLocation(input model.NewLocation) (*model.Location, error) {
	id := uuid.New()
	location := model.Location{
		ID:        fmt.Sprintf("%s", id),
		Name:      input.Name,
		Address:   input.Address,
		VisitedAt: input.VisitedAt,
	}
	err := lr.DB.Debug().Create(&location).Error
	return &location, err
}

func (lr *LocationRepository) getLocations() ([]*model.Location, error) {
	locations := []*model.Location{}
	err := lr.DB.Debug().Find(&locations).Error

	// クエリを作成
	query := lr.DB.Model(model.Location{}).Select("EXTRACT(YEAR FROM visited_at) as year").Distinct().Order("year desc")

	// クエリを実行
	var results []int
	query.Find(&results)

	return locations, err
}

func (lr *LocationRepository) getYears() ([]string, error) {
	years := []string{}
	err := lr.DB.Debug().Table("locations").Select("DISTINCT(visited_at)").Scan(&years).Error

	// クエリを作成
	query := lr.DB.Model(&model.Location{}).Select("year(visited_at) as year").Distinct().Order("year desc")

	// クエリを実行
	var results []int
	query.Find(&results)

	// 結果を表示
	for _, result := range results {
		log.Println(result)
	}
	return years, err
}
