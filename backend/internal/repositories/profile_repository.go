package repositories

import (
	"context"
	"profile-backend/config"
	"profile-backend/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var profileCollection *mongo.Collection = config.GetCollection("portfolio", "profiles")

func UpdateProfile(profile models.Profile) (*models.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": profile.ID}
	update := bson.M{
		"$set": bson.M{
			"name":         profile.Name,
			"bio":          profile.Bio,
			"contact_info": profile.ContactInfo,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := profileCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
