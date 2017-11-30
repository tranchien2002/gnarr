package structs

type Legistration struct {
	Name          string    `json:"name"`
	Sign          string    `json:"sign"`
	ID            string    `json:"id"`
	Enforcer      string    `json:"enforcer"`
	Baseon        string    `json:"baseon"`
	EffectiveDate string    `json:"effective_date"`
	PassDate      string    `json:"pass_date"`
	ChapterArray  []Chapter `json:"chapter_array"`
	ArticleArray  []Article `json:"article_array"`
}
