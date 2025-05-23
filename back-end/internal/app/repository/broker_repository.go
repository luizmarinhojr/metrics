package repository

import (
	"github.com/luizmarinhojr/metrics/internal/app/model"
	"gorm.io/gorm"
)

type BrokerRepository struct {
	db *gorm.DB
}

func newBrokerRepository(db *gorm.DB) *BrokerRepository {
	return &BrokerRepository{
		db: db,
	}
}

func (br *BrokerRepository) GetAll() *[]model.Broker {
	var brokers []model.Broker
	br.db.Find(&brokers)
	return &brokers
}

func (br *BrokerRepository) GetByName(name *string) *model.Broker {
	var broker model.Broker
	br.db.Where("nome = ?").First(&broker)
	return &broker
}
