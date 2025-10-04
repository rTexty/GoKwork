package types

// Subcategory представляет подкатегорию
type Subcategory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Category представляет категорию
type Category struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description,omitempty"`
	Subcategories []Subcategory `json:"subcategories,omitempty"`
}
