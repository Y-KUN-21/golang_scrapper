package models

type DetailAnime struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Studio        string   `json:"studio"`
	Tags          string   `json:"tags"`
	Year          string   `json:"year"`
	Season        string   `json:"Season"`
	Status        string   `json:"Status"`
	EpisodeNumber []string `json:"episodeNumber"`
	Episodes      []string `json:"episodes"`
}
