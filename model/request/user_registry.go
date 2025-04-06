package request

import "time"

type Country struct {
	Nation        string `json:"nation"`
	RegionOrState string `json:"region_or_state"`
	City          string `json:"city"`
}

type User struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"` // Consider using time.Time
	Country  Country   `json:"country"`
	Email    string    `json:"email"`
}
