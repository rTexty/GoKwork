package types

// User представляет пользователя
type User struct {
	ID                       string          `json:"id"`
	Username                 string          `json:"username"`
	Status                   string          `json:"status"`
	Fullname                 string          `json:"fullname"`
	ProfilePicture           string          `json:"profilepicture"`
	Description              string          `json:"description"`
	Slogan                   string          `json:"slogan"`
	Location                 string          `json:"location"`
	Rating                   float64         `json:"rating"`
	RatingCount              int             `json:"rating_count"`
	LevelDescription         string          `json:"level_description"`
	GoodReviews              int             `json:"good_reviews"`
	BadReviews               int             `json:"bad_reviews"`
	Online                   bool            `json:"online"`
	LiveDate                 int             `json:"live_date"`
	Cover                    string          `json:"cover"`
	CustomRequestMinBudget   int             `json:"custom_request_min_budget"`
	IsAllowCustomRequest     int             `json:"is_allow_custom_request"`
	OrderDonePercent         int             `json:"order_done_persent"`
	OrderDoneInTimePercent   int             `json:"order_done_intime_persent"`
	OrderDoneRepeatPercent   int             `json:"order_done_repeat_persent"`
	TimezoneID               int             `json:"timezoneId"`
	BlockedByUser            bool            `json:"blocked_by_user"`
	AllowedDialog            bool            `json:"allowedDialog"`
	AddTime                  int             `json:"addtime"`
	AchievementsList         []Achievement   `json:"achievments_list,omitempty"`
	CompletedOrdersCount     int             `json:"completed_orders_count"`
	Specialization           string          `json:"specialization,omitempty"`
	Profession               string          `json:"profession,omitempty"`
	KworksCount              int             `json:"kworks_count"`
	Kworks                   []KworkObject   `json:"kworks"`
	PortfolioList            interface{}     `json:"portfolio_list,omitempty"`
	Reviews                  []Review        `json:"reviews,omitempty"`
}
