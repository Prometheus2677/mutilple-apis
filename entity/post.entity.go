package entity

type Post struct {
	Id    uint64 `gorm:"primary_key:auto_increment" json: "id"`
	Title string `gorm:"not null" json:"title"`
	Body  string `gorm: "not null" json: "body"`
}
