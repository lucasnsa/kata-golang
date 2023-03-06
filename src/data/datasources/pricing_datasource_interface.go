package datasources

import (
	domain "pinpag/codingdojogolang/src/domain/entities"
	"time"
)

type PricingDataSourceInterface interface {
	FetchPricingFromDate(date time.Time, code string, codein string) (domain.Price, error)
}
