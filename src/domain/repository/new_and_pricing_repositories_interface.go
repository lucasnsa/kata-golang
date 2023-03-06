package domain

import (
	domain "pinpag/codingdojogolang/src/domain/entities"
	"time"
)

type NewAndPricingRepositoriesInterface interface {
	GetAllNewsAndPricingFromDate(date time.Time, code string, codein string) (domain.NewsAndPrice, error)
}
