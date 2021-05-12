package entities

type Article struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}
