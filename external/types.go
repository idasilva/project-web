package external

import "time"

type RequestNews struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Title     string    `gorm:"size:255;not null;unique" json:"title,omitempty" validator:"gte=10,lte=500"`
	Body      string    `gorm:"size:255;not null;unique" json:"body,omitempty" validator:"gte=10,lte=500"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" validator:"required"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at" validator:"required"`
	Data      string    `gorm:"size:255;not null;unique" json:"data,omitempty"`
}
type UrlImage struct {
	Image string `gorm:"size:255;not null;unique" json:"image,omitempty"`
}


var endpoint1 = "/find/all/documents"
var endpoint2 = "/register/documents"
var endpoint3= "/version"

