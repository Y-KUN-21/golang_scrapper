package models

type DetailAnime struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Tags          string   `json:"tags"`
	EpisodeNumber []string `json:"episodeNumber"`
	EpisodeURLs   []string `json:"episodeURLs"`
}
