package domain

import (
	domain "pinpag/codingdojogolang/domain/entities"
	"time"
)

type NewAndPricingRepositoriesInterface interface {
	GetAllNewsAndPricingFromDate(date time.Time, code string, codein string) (domain.NewsAndPrice, error)
}
