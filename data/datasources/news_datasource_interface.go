package datasources

import (
	domain "pinpag/codingdojogolang/domain/entities"
	"time"
)

type NewsDataSourceInterface interface {
	FetchLastestNewsFromDate(date time.Time, code string, codein string) ([]domain.NewsEntity, error)
}
