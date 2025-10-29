package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database


func ConnectDb() *mongo.Database{
	
	// loading the env file 
	err := godotenv.Load()
	if err != nil{
		log.Fatal("error loading the env file ", err)
	}
	//getting the uri for db and the dbname
	MongoUri := os.Getenv("MONGO_URI")
	DbName := os.Getenv("DB_NAME")

	clientOptions := options.Client().ApplyURI(MongoUri)

	// connecting to the Mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal("error connecting to the MongoDb database ", err)
	}
	fmt.Println("Successfully connected to the Mongodb database")

	DB = client.Database(DbName)

	return DB

}
