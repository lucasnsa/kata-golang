package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	domain "pinpag/codingdojogolang/domain/entities"
	"time"
)

type NewsDataSourceImpl struct {
	HttpClient http.Client
}

func (n *NewsDataSourceImpl) FetchLastestNewsFromDate(date time.Time, code string, codein string) ([]domain.NewsEntity, error) {
	time.Sleep(1 * time.Second)

	req, err := http.NewRequest("GET", "https://newsapi.org/v2/everything?", nil)
	if err != nil {
		fmt.Print(err)
	}

	q := req.URL.Query()
	q.Add("q", fmt.Sprintf("%s%s", code, codein))
	q.Add("from", date.Format("2006-01-02"))
	q.Add("sortBy", "publishedAt")
	q.Add("apiKey", "f60266f6f17f4a76b6d889ba48264aa8")
	q.Add("language", "en")
	q.Add("pageSize", "10")
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	response, responseErr := n.HttpClient.Do(req)

	if responseErr != nil {
		fmt.Println("ERROR:", err)
	}

	defer response.Body.Close()
	// Extrai conteúdo para salvar no banco de dados
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var newsResponse newsApiResponse
	json.Unmarshal([]byte(content), &newsResponse)

	return newsResponse.Articles, nil
}

type newsApiResponse struct {
	TotalResult string              `json:"totalResults"`
	Articles    []domain.NewsEntity `json:"articles"`
}
