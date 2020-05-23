package mangadex

type response struct {
	Manga   Manga              `json:"manga"`
	Chapter map[string]Chapter `json:"chapter"`
}
