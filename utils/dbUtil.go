package utils

import (
	. "awesomeProject/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const dbName = "db1"
const collectionName = "movies"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(os.Getenv("connectionstring"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connection successful")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection is ready", collection.Name())
}

func CreateOneMovie(movie *Movie) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie title added to the database", inserted.InsertedID)
}

func GetAllMovies() []primitive.M {
	movieCursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for movieCursor.Next(context.Background()) {
		var movie bson.M
		err := movieCursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	movieCursor.Close(context.Background())
	return movies
}
func GetMovieDetails(id string) Movie {
	fmt.Println("Fetching details of the movie ", id)
	movieD := collection.FindOne(context.Background(), bson.M{"uniqueid": id})
	var movieDetails Movie
	_ = movieD.Decode(&movieDetails)
	//if err != nil {
	//	log.Fatal(err)
	//	return movieDetails
	//}
	//_ = cursor.Decode(&movieDetails)
	//cursor.Close(context.Background())
	fmt.Println("Movie details ", movieDetails)
	return movieDetails
}
func UpdateMovieDetails(id string, updates *Movie) int64 {
	filter := bson.M{"uniqueid": id}
	update := bson.M{"$set": bson.M{"watched": updates.Watched, "moviename": updates.MovieName, "uniqueid": updates.UniqueId}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	fmt.Println("Updated records", result.MatchedCount)
	return result.ModifiedCount
}

func DeleteMovieRecord(id string) int64 {
	fmt.Println("Deleting movie record...", id)
	filter := bson.M{"uniqueid": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return result.DeletedCount
}
