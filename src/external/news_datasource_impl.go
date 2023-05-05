package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	domain "pinpag/codingdojogolang/src/domain/entities"
	"time"
)

type NewsDataSourceImpl struct {
	HttpClient http.Client
}

func (n *NewsDataSourceImpl) FetchLastestNewsFromDate(date time.Time, code string, codein string) ([]domain.News, error) {
	time.Sleep(2 * time.Second)

	req, err := http.NewRequest("GET", "https://newsapi.org/v2/everything?", nil)
	if err != nil {
		fmt.Print(err)
	}

	q := req.URL.Query()
	q.Add("q", fmt.Sprintf("%s%s", code, codein))
	q.Add("from", date.Format("2006-01-02"))
	q.Add("sortBy", "popularity")
	q.Add("apiKey", "enter_your_api_key")
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
	TotalResults string        `json:"totalResults"`
	Articles     []domain.News `json:"articles"`
}
