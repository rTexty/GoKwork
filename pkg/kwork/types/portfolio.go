package types

// PortfolioItem представляет элемент портфолио
type PortfolioItem struct {
	ID            int                      `json:"id"`
	Title         string                   `json:"title"`
	OrderID       interface{}              `json:"order_id"` // может быть int или string
	CategoryID    int                      `json:"category_id"`
	CategoryName  string                   `json:"category_name"`
	Type          string                   `json:"type"`
	Photo         string                   `json:"photo"`
	Video         string                   `json:"video"`
	Likes         int                      `json:"likes"`
	LikesDirty    int                      `json:"likes_dirty"`
	Views         int                      `json:"views"`
	ViewsDirty    int                      `json:"views_dirty"`
	CommentsCount int                      `json:"comments_count"`
	IsLiked       bool                     `json:"is_liked"`
	Images        []map[string]interface{} `json:"images,omitempty"`
	Videos        []map[string]interface{} `json:"videos,omitempty"`
	DuplicateFrom string                   `json:"duplicate_from"`
}
