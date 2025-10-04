package types

// Cover представляет обложку кворка
type Cover struct {
	Phone  string `json:"phone"`
	Tablet string `json:"tablet"`
}

// Worker представляет информацию о работнике
type Worker struct {
	ID             int     `json:"id"`
	Username       string  `json:"username"`
	Fullname       string  `json:"fullname"`
	ProfilePicture string  `json:"profilepicture"`
	Rating         float64 `json:"rating"`
	ReviewsCount   int     `json:"reviews_count"`
	RatingCount    int     `json:"rating_count"`
	IsOnline       bool    `json:"is_online"`
}

// Activity представляет активность кворка
type Activity struct {
	Views  int `json:"views"`
	Orders int `json:"orders"`
	Earned int `json:"earned"`
}

// KworkObject представляет объект кворка
type KworkObject struct {
	ID             int      `json:"id"`
	CategoryID     int      `json:"category_id"`
	CategoryName   string   `json:"category_name"`
	StatusID       int      `json:"status_id"`
	StatusName     string   `json:"status_name"`
	Title          string   `json:"title"`
	URL            string   `json:"url"`
	ImageURL       string   `json:"image_url"`
	Cover          *Cover   `json:"cover,omitempty"`
	Price          int      `json:"price"`
	IsPriceFrom    bool     `json:"is_price_from"`
	IsFrom         bool     `json:"is_from"`
	Photo          string   `json:"photo"`
	IsBest         bool     `json:"is_best"`
	IsHidden       bool     `json:"is_hidden"`
	IsFavorite     bool     `json:"is_favorite"`
	Lang           string   `json:"lang"`
	Worker         *Worker  `json:"worker,omitempty"`
	Activity       *Activity `json:"activity,omitempty"`
	EditsList      []string `json:"edits_list,omitempty"`
	ProfileSort    int      `json:"profile_sort"`
	IsSubscription bool     `json:"isSubscription"`
	Badges         []interface{} `json:"badges,omitempty"`
}
