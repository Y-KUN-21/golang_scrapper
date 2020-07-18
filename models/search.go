package models

type SearchAnime struct {
	Name     string `json:"name"`
	Year     string `json:"year"`
	Season   string `json:"season"`
	Status   string `json:"status"`
	ImageURL string `json:"imageurl"`
	Url      string `json:"url"`
}
