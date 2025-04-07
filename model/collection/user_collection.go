package collection

import "time"

type Country struct {
	Nation        string `bson:"nation" json:"nation"`
	RegionOrState string `bson:"region_or_state" json:"region_or_state"`
	City          string `bson:"city" json:"city"`
}

type User struct {
	Username     string    `bson:"username" json:"username"`
	Password     string    `bson:"password" json:"password"`
	Birthday     time.Time `bson:"birthday" json:"birthday"` // Consider using time.Time
	Country      Country   `bson:"country" json:"country"`
	Email        string    `bson:"email" json:"email"`
	ClientID     string    `bson:"clientId" json:"clientId"`
	ClientSecret string    `bson:"clientSecret" json:"clientSecret"`
}
