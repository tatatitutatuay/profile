package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profile struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Bio         string             `json:"bio" bson:"bio"`
	ContactInfo string             `json:"contact_info" bson:"contact_info"`
}
