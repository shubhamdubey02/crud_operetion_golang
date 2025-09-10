package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CleanupUsersCollection(ctx context.Context, collection *mongo.Collection) {
	unsetFields := bson.M{
		"About":    "",
		"Address":  "",
		"Branch":   "",
		"College":  "",
		"Language": "",
		"Phone":    "",
		"Projects": "",
		"Role":     "",
	}

	update := bson.M{"$unset": unsetFields}

	result, err := collection.UpdateMany(ctx, bson.M{}, update)
	if err != nil {
		log.Fatal("UpdateMany error:", err)
	}
	if result.ModifiedCount > 0 {
		fmt.Printf("Cleanup complete: %d documents updated\n", result.ModifiedCount)
	} else {
		fmt.Println("Db connected to MongoDB")
	}
}
