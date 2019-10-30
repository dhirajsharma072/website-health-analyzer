package models

import (
	"time"
)

//Website represents a music site
type Website struct {
	ID      string    `bson:"uuid,omitempty" json:"id,omitempty"`
	URL       string    `bson:"url,omitempty" json:"url,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	IsHealthy bool      `bson:"is_healthy,omitempty" json:"isHealthy,omitempty"`
}

type WebsitePatch struct {
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	IsHealthy bool      `bson:"is_healthy,omitempty" json:"isHealthy,omitempty"`
}

//Websites is an array of Website
type Websites []Website
