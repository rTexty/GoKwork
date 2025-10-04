package types

// KworkMinObject представляет минимальную информацию о кворке
type KworkMinObject struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Active int    `json:"active"`
	Feat   bool   `json:"feat"`
}

// Writer представляет автора отзыва
type Writer struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profilepicture"`
}

// Review представляет отзыв
type Review struct {
	ID        int             `json:"id"`
	TimeAdded int             `json:"time_added"`
	Text      string          `json:"text"`
	AutoMode  string          `json:"auto_mode,omitempty"`
	Good      bool            `json:"good"`
	Bad       bool            `json:"bad"`
	Kwork     *KworkMinObject `json:"kwork,omitempty"`
	Writer    *Writer         `json:"writer,omitempty"`
}
