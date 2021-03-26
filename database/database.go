package database

import (
	"context"
	"log"
	"phrases-server/config"
	"phrases-server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// Connect establish a connection to models
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the models
	db = client.Database(config.Config.DatabaseName)
}

func GetReviewPhrases(frequencySortOrder SortOrder) []models.ReviewPhrase {
	const CollectionName = "review_phrases"

	options := options.Find().SetSort(bson.D{{"frequency", frequencySortOrder}})

	cur, err := db.Collection(CollectionName).Find(context.Background(), bson.D{{}}, options)
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.ReviewPhrase
	var elem models.ReviewPhrase
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}
