package types

// Connects представляет информацию о коннектах
type Connects struct {
	AllConnects    int `json:"all_connects"`
	ActiveConnects int `json:"active_connects"`
	UpdateTime     int `json:"update_time"`
}
