package types

// Project представляет проект на бирже
type Project struct {
	ID                  int           `json:"id"`
	UserID              int           `json:"user_id"`
	Username            string        `json:"username"`
	ProfilePicture      string        `json:"profile_picture"`
	Price               int           `json:"price"`
	Title               string        `json:"title"`
	Description         string        `json:"description"`
	Offers              int           `json:"offers"`
	TimeLeft            int           `json:"time_left"`
	ParentCategoryID    int           `json:"parent_category_id"`
	CategoryID          int           `json:"category_id"`
	DateConfirm         int           `json:"date_confirm"`
	CategoryBasePrice   int           `json:"category_base_price"`
	UserProjectsCount   int           `json:"user_projects_count"`
	UserHiredPercent    int           `json:"user_hired_percent"`
	AchievementsList    []Achievement `json:"achievements_list,omitempty"`
	IsViewed            bool          `json:"is_viewed"`
	AlreadyWork         int           `json:"already_work"`
	AllowHigherPrice    bool          `json:"allow_higher_price"`
	PossiblePriceLimit  int           `json:"possible_price_limit"`
}
