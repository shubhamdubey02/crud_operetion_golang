package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetISTTime() string {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return time.Now().UTC().Format("2006-01-02 T 03:04 PM")
	}
	return time.Now().In(loc).Format("2006-01-02 T 03:04 PM")
}

type Schema struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Age       int                `bson:"age" json:"age"`
	Email     string             `bson:"email" json:"email"`
	CreatedAt string             `bson:"created_at" json:"created_at"`
	UpdatedAt string             `bson:"updated_at" json:"updated_at"`
}
