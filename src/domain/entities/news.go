package domain

type News struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
