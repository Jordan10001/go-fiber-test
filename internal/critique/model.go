package critique

import "go.mongodb.org/mongo-driver/bson/primitive"

// Critique represents a single user feedback entry.
type Critique struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}