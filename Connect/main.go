package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" //Thanks to joho!
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email string             `json:"email"`
}

var user []User

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI") //CONFIGURE YOUR .env WITH YOUR URI HAVING "MONGO_URI" AS KEY VALUE
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("GolangDb").Collection("GolangCollection")
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var newUser User
		if err = cursor.Decode(&newUser); err != nil {
			log.Fatal(err)
		}
		user = append(user, newUser)
	}

	for _, x := range user {
		fmt.Println(x)
	}

	/* SINGLE GET

	var podcast bson.M
	if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcast)

	*/
}
