package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schema struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Age       int                `bson:"age" json:"age"`
	Email     string             `bson:"email" json:"email"`
	College   string             `bson:"college" json:"college"`
	Branch    string             `bson:"branch" json:"branch"`
	Phone     string             `bson:"phone" json:"phone"`
	Address   string             `bson:"address" json:"address"`
	Language  string             `bson:"language" json:"language"`
	Projects  string             `bson:"projects" json:"projects"`
	About     string             `bson:"about" json:"about"`
	Role      string             `bson:"role" json:"role"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
