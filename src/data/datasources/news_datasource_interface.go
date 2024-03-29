package datasources

import (
	domain "pinpag/codingdojogolang/src/domain/entities"
	"time"
)

type NewsDataSourceInterface interface {
	FetchLastestNewsFromDate(date time.Time, code string, codein string) ([]domain.News, error)
}
