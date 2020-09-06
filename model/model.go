package model

type MyJsonName struct {
	List []struct {
		Author      string        `json:"author"`
		CurrentVote string        `json:"current_vote"`
		Defid       int64         `json:"defid"`
		Definition  string        `json:"definition"`
		Example     string        `json:"example"`
		Permalink   string        `json:"permalink"`
		SoundUrls   []interface{} `json:"sound_urls"`
		ThumbsDown  int64         `json:"thumbs_down"`
		ThumbsUp    int64         `json:"thumbs_up"`
		Word        string        `json:"word"`
		WrittenOn   string        `json:"written_on"`
	} `json:"list"`
}
