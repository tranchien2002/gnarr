package structs

type Topic struct {
	FirstArticle int64  `json:"first_article"`
	LastArticle  int64  `json:"last_article"`
	Header       string `json:"topic_header"`
}
