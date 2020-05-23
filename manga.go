package mangadex

type Manga struct {
	CoverURL    string    `json:"cover_url"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	Artist      string    `json:"artist"`
	Author      string    `json:"author"`
	Status      int64     `json:"status"`
	Genres      []int64   `json:"genres"`
	Language    string    `json:"lang_flag"`
	Chapters    []Chapter `json:"chapters"`
}
