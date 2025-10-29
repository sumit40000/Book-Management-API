package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sumit40000/book-management-api/internal/config"
	"github.com/sumit40000/book-management-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Collection = config.DB.Collection("Books")


// helper functions for controllers

// creating a book
func createBook(book models.Book) interface{}{
	created, err := Collection.InsertOne(context.Background(), book)
	if err != nil{
		log.Fatal("ERROR! inserting the book in thte database", err)
	}

	return created.InsertedID

}

// updating a book using the id 
func updateBook(bookId string) int64{
	id, err := primitive.ObjectIDFromHex(bookId)
	if err != nil{
		log.Fatal("ERROR! converting the string id to objectid", err)
	}
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"read":true}}
	updated, err := Collection.UpdateByID(context.Background(),filter, update)
	if err != nil {
		log.Fatal("ERROR! updating the book try again..", err)
	}
	 return updated.ModifiedCount
}

// finding a book using it's id
func findOneBook(bookId string) models.Book {
	var book models.Book
	id , err := primitive.ObjectIDFromHex(bookId)
	if err != nil{
		log.Fatal("ERROR! converting the string id to objectid", err)
	}
	filter := bson.M{"_id":id}
	err = Collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil{
		log.Fatal("ERROR! decoding the data ",err)
	}
	return book
}

// getting all books from database
func getAllBooks() []primitive.M{
	cursor, err := Collection.Find(context.Background(), bson.M{})
	if err!= nil{
		log.Fatal("ERROR! getting all the books try again", err)
	}
	defer cursor.Close(context.Background())
	var books []bson.M
	for cursor.Next(context.Background()){
		var book bson.M
		err := cursor.Decode(&book)
		if err!= nil{
			log.Fatal("error decoding", err)
		}
		books = append(books, book)

	}
	
	return books
}

// deleting a book by it's id 
func deleteOneBook(bookId string) int64{
	id , err := primitive.ObjectIDFromHex(bookId)
	if err != nil{
		log.Fatal("error converting to objectid",err)
	}
	filter := bson.M{"_id":id}
	result, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil{
		log.Fatal("error deleting the book ",err)
	}
	return result.DeletedCount
}

// deleting all books from database
func deleteAllBooks() int64{
	result, err := Collection.DeleteMany(context.Background(),bson.M{})
	if err != nil{
		log.Fatal("error deleting all books from database",err)
	}
	return result.DeletedCount
}

// controllers 

// creating a book
func CreateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods","POST")

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil{
		log.Fatal("error decoding the data ", err)
	}
	created := createBook(book)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "book created successfully",
		"insertID" : created,
	})
}

// updating a book 
func UpdateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updatecount := updateBook(params["id"])
	json.NewEncoder(w).Encode(map[string]any{
		"message": "book updated successfully",
		"insertID" : updatecount,
	})
}

// getting a book from its id 
func FindOneBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	foundBook := findOneBook(params["id"])
	err:= json.NewEncoder(w).Encode(foundBook)
	if err != nil{
		log.Fatal("error returning the book try again..", err)
	}
}

// getting all books
func GetAllBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	allbooks := getAllBooks()

	err:= json.NewEncoder(w).Encode(allbooks)
	if err != nil{
		log.Fatal("error returning the book try again..", err)
	}
}


// deleting a book 
func DeleteOneBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	params:= mux.Vars(r)
	deletecount := deleteOneBook(params["id"])
	json.NewEncoder(w).Encode(map[string]any{
		"message": "book deleted successfully",
		"deletedNo" : deletecount,
	})

}

// deleting all book
func DeleteAllBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	deletedcount := deleteAllBooks()
	json.NewEncoder(w).Encode(map[string]any{
		"message": "All books deleted successfully",
		"insertID" : deletedcount,
	})
}