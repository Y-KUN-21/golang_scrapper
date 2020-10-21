package models

type SearchAnime struct {
	Name   string `json:"name"`
	Year   string `json:"year"`
	Season string `json:"season"`
	Status string `json:"status"`
	Cover  string `json:"cover"`
	Url    string `json:"url"`
}
